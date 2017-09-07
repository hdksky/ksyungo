package slb

import (
	"github.com/hdksky/ksyungo/common"
	"github.com/hdksky/ksyungo/util"

	"errors"
	ip "github.com/1851616111/util/validator/ipv4"
	pt "github.com/1851616111/util/validator/port"
	"regexp"
)

type RegisterInstancesWithListenerArgs struct {
	ListenerId     string
	RealServerIp   string
	RealServerPort int
	RealServerType string
	Weight         uint8
}

func (a *RegisterInstancesWithListenerArgs) validate() error {
	if len(a.ListenerId) == 0 {
		return util.ParamNotFoundErr("ListenerId")
	}

	if _, err := regexp.MatchString(uuidRegexp, a.ListenerId); err != nil {
		return util.ParamInvalid("ListenerId")
	}

	if err := ip.Validate(a.RealServerIp); err != nil {
		return err
	}

	if err := pt.ValidatePortNum(a.RealServerPort); err != nil {
		return err
	}
	//
	//if a.RealServerType != "host" {
	//	a.RealServerType = "host"
	//}

	if a.Weight == 0 {
		a.Weight = 1
	}

	return nil
}

type RealServerType struct {
	RegisterId      string
	ListenerId      string
	RealServerIp    string
	RealServerPort  int
	RealServerType  string
	RealServerState string
	Weight          int
}

type RegisterInstancesWithListenerResponse struct {
	common.Response
	RealServerType
}

// RegisterInstancesWithListener Register Instances With Listener
// You can read doc at https://docs.ksyun.com/read/latest/55/_book/Action/RegisterInstancesWithListener.html
func (c *Client) RegisterInstancesWithListener(args *RegisterInstancesWithListenerArgs) (*RegisterInstancesWithListenerResponse, error) {
	if args == nil {
		return nil, errors.New("registe instance with listener with no args")
	}

	if err := args.validate(); err != nil {
		return nil, err
	}

	response := RegisterInstancesWithListenerResponse{}
	err := c.Invoke("RegisterInstancesWithListener", args, &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}

type ModifyInstancesWithListenerArgs struct {
	RegisterId     string
	RealServerPort int
	Weight         int
}

type ModifyInstancesWithListenerResponse struct {
	common.Response
	RealServerType
}

// ModifyInstancesWithListener Register Instances With Listener
// You can read doc at https://docs.ksyun.com/read/latest/55/_book/Action/ModifyInstancesWithListener.html
func (c *Client) ModifyInstancesWithListener(args *ModifyInstancesWithListenerArgs) (*ModifyInstancesWithListenerResponse, error) {
	response := ModifyInstancesWithListenerResponse{}
	err := c.Invoke("ModifyInstancesWithListener", args, &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}

type DeregisterInstancesFromListenerArgs struct {
	RegisterId string
}
type DeregisterInstancesFromListenerResponse struct {
	common.Response
	Return bool
}

// DeregisterInstancesFromListener Register Instances With Listener
// You can read doc at https://docs.ksyun.com/read/latest/55/_book/Action/DeregisterInstancesFromListener.html
func (c *Client) DeregisterInstancesFromListener(args *DeregisterInstancesFromListenerArgs) (bool, error) {
	response := DeregisterInstancesFromListenerResponse{}
	if err := c.Invoke("DeregisterInstancesFromListener", args, &response); err != nil {
		return false, err
	}

	return response.Return, nil
}

type DescribeInstancesWithListenerArgs struct {
	RegisterId []string
	Filter     []KV
}

type DescribeInstancesWithListenerResponse struct {
	common.Response
	RealServerSet struct {
		Item []RealServerType
	}
}

// DescribeInstancesWithListener Register Instances With Listener
// You can read doc at https://docs.ksyun.com/read/latest/55/_book/Action/DescribeInstancesWithListener.html
func (c *Client) DescribeInstancesWithListener(args *DescribeInstancesWithListenerArgs) ([]RealServerType, error) {
	response := DescribeInstancesWithListenerResponse{}
	err := c.Invoke("DescribeInstancesWithListener", args, &response)
	if err == nil {
		return response.RealServerSet.Item, nil
	}
	return nil, err
}
