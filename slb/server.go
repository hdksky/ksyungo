package slb

import "github.com/hdksky/ksyungo/common"

type RegisterInstancesWithListenerArgs struct {
	ListenerId     string
	RealServerIp   string
	RealServerPort int
	RealServerType string
	Weight         int
}

type RealServerType struct {
	RegisterId      string
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

type DeregisterInstancesFromListenerResponse struct {
	common.Response
	Return bool
}

// DeregisterInstancesFromListener Register Instances With Listener
// You can read doc at https://docs.ksyun.com/read/latest/55/_book/Action/DeregisterInstancesFromListener.html
func (c *Client) DeregisterInstancesFromListener(registerId string) (bool, error) {
	response := DeregisterInstancesFromListenerResponse{}
	err := c.Invoke("DeregisterInstancesFromListener", registerId, &response)
	if err == nil {
		return response.Return, nil
	}
	return false, err
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
