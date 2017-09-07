package slb

import (
	"errors"
	"github.com/hdksky/ksyungo/common"
	"github.com/hdksky/ksyungo/util"
	"regexp"
)

var uuidRegexp string = `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`
var ListenerMethod_RoundRobin string = "RoundRobin"
var ListenerMethod_LeastConnections string = "LeastConnections"
var ListenerState_Start = "start"
var ListenerState_Stop = "stop"

var CookieType_ImplantCookie = "ImplantCookie"
var CookieType_RewriteCookie = "RewriteCookie"

type CreateListenersArgs struct {
	LoadBalancerId           string
	ListenerState            string
	ListenerName             string
	ListenerProtocol         string
	CertificateId            string
	ListenerPort             uint16
	Method                   string
	SessionState             string
	SessionPersistencePeriod int
	CookieType               string
	CookieName               string
}

func (a *CreateListenersArgs) validate() error {
	if len(a.LoadBalancerId) == 0 {
		return util.ParamNotFoundErr("LoadBalancerId")
	}

	if _, err := regexp.MatchString(uuidRegexp, a.LoadBalancerId); err != nil {
		return util.ParamInvalid("LoadBalancerId")
	}

	if len(a.ListenerState) == 0 {
		a.ListenerState = "start"
	}

	if a.ListenerState != ListenerState_Start && a.ListenerState != ListenerState_Stop {
		return util.ParamNotInValidList("ListenerState", ListenerState_Start, ListenerState_Stop)
	}

	if len(a.ListenerProtocol) == 0 {
		return util.ParamNotFoundErr("ListenerProtocol")
	}

	if a.ListenerProtocol != "TCP" && a.ListenerProtocol != "HTTP" && a.ListenerProtocol != "HTTPS" {
		return util.ParamNotInValidList("ListenerProtocol", "TCP", "HTTP", "HTTPS")
	}

	if a.ListenerProtocol == "HTTPS" && len(a.CertificateId) == 0 {
		return util.ParamNotFoundErr("CertificateId")
	}

	if a.ListenerPort < 1 || a.ListenerPort > 65535 {
		return util.ParamValueInvalid("ListenerPort", 1, 65535)
	}

	if len(a.Method) == 0 {
		return util.ParamNotFoundErr("Method")
	}

	if a.Method != ListenerMethod_RoundRobin && a.Method != ListenerMethod_LeastConnections {
		return util.ParamNotInValidList("Method", ListenerMethod_RoundRobin, ListenerMethod_LeastConnections)
	}

	if len(a.SessionState) == 0 {
		return util.ParamNotFoundErr("SessionState")
	}

	if a.SessionState != "start" && a.SessionState != "stop" {
		return util.ParamNotInValidList("SessionState", "start", "stop")
	}

	if a.SessionPersistencePeriod == 0 {
		a.SessionPersistencePeriod = 3600
	}

	if a.SessionPersistencePeriod > 86400 || a.SessionPersistencePeriod < 1 {
		return util.ParamValueInvalid("SessionPersistencePeriod", 1, 86400)
	}

	if a.ListenerProtocol == "HTTP" || a.ListenerProtocol == "HTTPS" {
		if a.CookieType != CookieType_ImplantCookie && a.CookieType != CookieType_RewriteCookie && a.CookieType != "" {
			return util.ParamNotInValidList("CookieType", CookieType_ImplantCookie, CookieType_RewriteCookie, "")
		}

		if a.CookieType == CookieType_RewriteCookie {
			if len(a.CookieName) == 0 {
				return util.ParamNotFoundErr("CookieName")
			}
		}
	}

	return nil
}

type SessionType struct {
	SessionState             string
	SessionPersistencePeriod int
}

type ListenerType struct {
	CreateTime       string
	LoadBalancerId   string
	ListenerName     string
	ListenerId       string
	ListenerState    string
	CertificateId    string
	ListenerProtocol string
	ListenerPort     string
	Method           string
	HealthCheck      HealthCheckType
	Session          SessionType
	RealServer       struct {
		Item []RealServerType
	}
}

type CreateListenersResponse struct {
	common.Response
	ListenerType
}

// CreateListeners Create listener
// Only TCP protocal are allowed on type=internal LoadBalancer
// You can read doc at https://docs.ksyun.com/read/latest/55/_book/Action/CreateListeners.html
func (c *Client) CreateListeners(args *CreateListenersArgs) (*ListenerType, error) {
	if args == nil {
		return nil, errors.New("create listener args not found")
	}

	if err := args.validate(); err != nil {
		return nil, err
	}

	lbs, err := c.DescribeLoadBalancers(&DescribeLoadBalancersArgs{
		LoadBalancerId: []string{args.LoadBalancerId},
	})
	if err != nil {
		return nil, err
	}

	if len(lbs) == 0 {
		return nil, errors.New("LoadBalancer not found")
	}

	if lbs[0].Type == LoadBalancerType_Internal && args.ListenerProtocol != "TCP" {
		return nil, errors.New("only tcp protocol allowed on internal loadBalancer")
	}

	response := CreateListenersResponse{}
	err = c.Invoke("CreateListeners", args, &response)
	if err == nil {
		return &response.ListenerType, nil
	}
	return nil, err
}

type ModifyListenersArgs struct {
	ListenerId               string
	ListenerName             string
	ListenerState            string
	SessionState             string
	Method                   string
	SessionPersistencePeriod int
	CookieType               string
	CookieName               string
}

type ModifyListenersResponse struct {
	common.Response
	ListenerType
}

// ModifyListeners modify listener
// You can read doc at https://docs.ksyun.com/read/latest/55/_book/Action/ModifyListeners.html
func (c *Client) ModifyListeners(args *ModifyListenersArgs) (*ModifyListenersResponse, error) {
	response := ModifyListenersResponse{}
	err := c.Invoke("ModifyListeners", args, &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}

type DeleteListenerArgs struct {
	ListenerId string
}
type DeleteListenersResponse struct {
	common.Response
	Return bool
}

// DeleteListeners delete listener
// You can read doc at https://docs.ksyun.com/read/latest/55/_book/Action/DeleteListeners.html
func (c *Client) DeleteListeners(args *DeleteListenerArgs) (bool, error) {
	response := DeleteListenersResponse{}
	err := c.Invoke("DeleteListeners", args, &response)
	if err == nil {
		return response.Return, nil
	}
	return false, err
}

type DescribeListenersArgs struct {
	ListenerId []string
	Filter     []KV
}

type DescribeListenersResponse struct {
	common.Response
	ListenerSet struct {
		Item []ListenerType
	}
}

// DescribeListeners delete listener
// You can read doc at https://docs.ksyun.com/read/latest/55/_book/Action/DescribeListeners.html
func (c *Client) DescribeListeners(args *DescribeListenersArgs) ([]ListenerType, error) {
	response := DescribeListenersResponse{}
	err := c.Invoke("DescribeListeners", args, &response)
	if err == nil {
		return response.ListenerSet.Item, nil
	}
	return nil, err
}
