package slb

import (
	"time"

	"github.com/hdksky/ksyungo/common"
)

type CreateLoadBalancerArgs struct {
	VpcId            string
	LoadBalancerName string
	Type             string
	SubnetId         string
}

type LoadBalancerDescription struct {
	CreateTime        time.Time
	LoadBalancerName  string
	VpcId             string
	LoadBalancerId    string
	Type              string
	SubnetId          string
	PublicIp          string
	State             string
	LoadBalancerState string
}

type CreateLoadBalancerResponse struct {
	common.Response
	LoadBalancerDescription
}

// CreateLoadBalancer create load balancer
// You can read doc at https://docs.ksyun.com/read/latest/55/_book/Action/CreateLoadBalancer.html
func (c *Client) CreateLoadBalancer(args *CreateLoadBalancerArgs) ([]LoadBalancerDescription, error) {
	response := CreateLoadBalancerResponse{}
	err := c.Invoke("CreateLoadBalancer", args, &response)
	if err == nil {
		return []LoadBalancerDescription{{
			CreateTime:        response.CreateTime,
			LoadBalancerName:  response.LoadBalancerName,
			VpcId:             response.VpcId,
			LoadBalancerId:    response.LoadBalancerId,
			Type:              response.Type,
			SubnetId:          response.SubnetId,
			PublicIp:          response.PublicIp,
			State:             response.State,
			LoadBalancerState: response.LoadBalancerState,
		}}, nil
	}
	return nil, err
}

type DeleteLoadBalancerResponse struct {
	common.Response
	Return bool
}

// DeleteLoadBalancer delete load balancer
// You can read doc at https://docs.ksyun.com/read/latest/55/_book/Action/CreateLoadBalancer.html
func (c *Client) DeleteLoadBalancer(loadBalancerId string) (bool, error) {
	response := DeleteLoadBalancerResponse{}
	err := c.Invoke("DeleteLoadBalancer", loadBalancerId, &response)
	if err == nil {
		return response.Return, nil
	}
	return false, err
}

type ModifyLoadBalancerArgs struct {
	LoadBalancerId    string
	LoadBalancerName  string
	LoadBalancerState string
}

type ModifyLoadBalancerResponse struct {
	common.Response
	LoadBalancerDescription
}

// ModifyLoadBalancer modify load balancer
// You can read doc at https://docs.ksyun.com/read/latest/55/_book/Action/ModifyLoadBalancer.html
func (c *Client) ModifyLoadBalancer(args *ModifyLoadBalancerArgs) (*ModifyLoadBalancerResponse, error) {
	response := ModifyLoadBalancerResponse{}
	err := c.Invoke("ModifyLoadBalancer", args, &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}

type KV struct {
	Name  string
	Value string
}

type DescribeLoadBalancersArgs struct {
	LoadBalancerId []string
	State          string
	Filter         []KV
}

type LoadBalancerDescriptionResponse struct {
	common.Response
	LoadBalancerDescriptions struct {
		Item []LoadBalancerDescription
	}
}

// DescribeLoadBalancers describe load balancer
// You can read doc at https://docs.ksyun.com/read/latest/55/_book/Action/DescribeLoadBalancers.html
func (c *Client) DescribeLoadBalancers(args *DescribeLoadBalancersArgs) ([]LoadBalancerDescription, error) {
	response := LoadBalancerDescriptionResponse{}
	err := c.Invoke("DescribeLoadBalancers", args, &response)
	if err == nil {
		return response.LoadBalancerDescriptions.Item, nil
	}
	return nil, err
}
