package slb

import "github.com/hdksky/ksyungo/common"

type CreateListenersArgs struct {
	LoadBalancerId           string
	ListenerState            string
	ListenerName             string
	ListenerProtocol         string
	CertificateId            string
	ListenerPort             string
	Method                   string
	SessionState             string
	SessionPersistencePeriod int
	CookieType               string
	CookieName               string
}

type SessionType struct {
	SessionState           string
	CookieExpirationPeriod int
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
	Listener ListenerType
}

// CreateListeners Create listener
// You can read doc at https://docs.ksyun.com/read/latest/55/_book/Action/CreateListeners.html
func (c *Client) CreateListeners(args *CreateListenersArgs) (*ListenerType, error) {
	response := CreateListenersResponse{}
	err := c.Invoke("CreateListeners", args, &response)
	if err == nil {
		return &response.Listener, nil
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

type DeleteListenersResponse struct {
	common.Response
	Return bool
}

// DeleteListeners delete listener
// You can read doc at https://docs.ksyun.com/read/latest/55/_book/Action/DeleteListeners.html
func (c *Client) DeleteListeners(listenerId string) (bool, error) {
	response := DeleteListenersResponse{}
	err := c.Invoke("DeleteListeners", listenerId, &response)
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
