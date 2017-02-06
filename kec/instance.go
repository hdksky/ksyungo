package kec

import (
	"github.com/hdksky/ksyungo/common"
	"github.com/hdksky/ksyungo/util"
)

type KV struct {
	Name  string
	Value string
}

type DescribeInstancesArgs struct {
	MaxResults int
	Marker     int
	InstanceId string
	Filters    []KV
	Sort       []KV
	Search     string
}

type InstanceConfigureType struct {
	VCPU         int
	MemoryGb     int
	DataDiskType string
	DataDiskGb   int
}

type SecurityGroupGroup struct {
	SecurityGroupGroupId string
}

type NetworkInterfaceSetType struct {
	NetworkInterfaceId   string
	NetworkInterfaceType string
	MacAddress           string
	SubnetId             string
	VpcId                string
	PrivateIpAddress     string
	SecurityGroupSet     struct {
		Item []SecurityGroupGroup
	}
	DNS1 string
	DNS2 string
}

type Instance struct {
	InstanceId        string
	InstanceName      string
	InstanceConfigure InstanceConfigureType
	ImageId           string
	InstanceType      string
	InstanceState     struct {
		Name string
	}
	SubnetId         string
	PrivateIpAddress string
	Monitoring       struct {
		State string
	}
	SriovNetSupport     bool
	CreationDate        util.ISO6801Time
	NetworkInterfaceSet struct {
		Item []NetworkInterfaceSetType
	}
}

type DescribeInstancesResponse struct {
	common.Response
	Marker        int
	InstanceCount int
	InstancesSet  struct {
		Item []Instance
	}
}

// DescribeInstances describes instances
// You can read doc at https://docs.ksyun.com/read/latest/52/_book/oaDescribeInstances.html
func (c *Client) DescribeInstances(args *DescribeInstancesArgs) (*DescribeInstancesResponse, error) {
	response := DescribeInstancesResponse{}
	err := c.Invoke("DescribeInstances", args, &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}

type RunInstancesArgs struct {
	ImageId            string
	InstanceType       string
	DataDiskGb         string
	MaxCount           int
	MinCount           int
	SubnetId           string
	InstancePassword   string
	ChargeType         string
	PurchaseTime       int
	SecurityGroupId    string
	PrivateIpAddress   string
	InstanceName       string
	InstanceNameSuffix string
	SriovNetSupport    bool
}

type RunInstancesResponse struct {
	common.Response
	InstancesSet struct {
		Item []Instance
	}
}

// RunInstances create instances
// you can read doc at https://docs.ksyun.com/read/latest/52/_book/oaRunInstances.html
func (c *Client) RunInstances(args *RunInstancesArgs) ([]Instance, error) {
	response := RunInstancesResponse{}
	err := c.Invoke("RunInstances", args, &response)
	if err == nil {
		return response.InstancesSet.Item, nil
	}
	return nil, err
}

type TerminateInstancesArgs struct {
	InstanceId []string
}

type InstanceStateChange struct {
	InstanceId string
	Return     string
}

type InstancesStateChangeResponse struct {
	common.Response
	InstancesSet struct {
		Item []InstanceStateChange
	}
}

// TerminateInstances Terminate Instances
// you can read doc at https://docs.ksyun.com/read/latest/52/_book/oaTerminateInstances.html
// 约束条件1：在销毁实例时，需要先解绑实例已挂载的EBS数据盘、SLB、EIP资源
func (c *Client) TerminateInstances(args *TerminateInstancesArgs) ([]InstanceStateChange, error) {
	response := InstancesStateChangeResponse{}
	err := c.Invoke("TerminateInstances", args, &response)
	if err == nil {
		return response.InstancesSet.Item, nil
	}
	return nil, err
}

// StartInstances Start Instances
// you can read doc at https://docs.ksyun.com/read/latest/52/_book/oaStartInstances.html
func (c *Client) StartInstances(InstanceIds []string) ([]InstanceStateChange, error) {
	response := InstancesStateChangeResponse{}
	err := c.Invoke("StartInstances", InstanceIds, &response)
	if err == nil {
		return response.InstancesSet.Item, nil
	}
	return nil, err
}

// StopInstances Stop Instances
// you can read doc at https://docs.ksyun.com/read/latest/52/_book/oaStopInstances.html
func (c *Client) StopInstances(InstanceIds []string) ([]InstanceStateChange, error) {
	response := InstancesStateChangeResponse{}
	err := c.Invoke("StopInstances", InstanceIds, &response)
	if err == nil {
		return response.InstancesSet.Item, nil
	}
	return nil, err
}

// RebootInstances Reboot Instances
// you can read doc at https://docs.ksyun.com/read/latest/52/_book/oaRebootInstances.html
func (c *Client) RebootInstances(InstanceIds []string) ([]InstanceStateChange, error) {
	response := InstancesStateChangeResponse{}
	err := c.Invoke("RebootInstances", InstanceIds, &response)
	if err == nil {
		return response.InstancesSet.Item, nil
	}
	return nil, err
}

type ModifyInstanceAttributeArgs struct {
	InstanceId       string
	InstanceName     string
	InstancePassword string
}

type ModifyInstanceAttributeResponse struct {
	common.Response
	Return string
}

// ModifyInstanceAttribute Modify Instance Attribute
func (c *Client) ModifyInstanceAttribute(args *ModifyInstanceAttributeArgs) (string, error) {
	response := ModifyInstanceAttributeResponse{}
	err := c.Invoke("ModifyInstanceAttribute", args, &response)
	if err == nil {
		return response.Return, nil
	}
	return "", err
}

type ModifyInstanceTypeArgs struct {
	InstanceId   string
	InstanceType string
	DataDiskGb   int
}

type ModifyInstanceTypeResponse struct {
	common.Response
	Return bool
}

// ModifyInstanceType 升级实例套餐类型
// You can read doc at https://docs.ksyun.com/read/latest/52/_book/oaModifyInstanceType.html
func (c *Client) ModifyInstanceType(args *ModifyInstanceTypeArgs) (bool, error) {
	response := ModifyInstanceTypeResponse{}
	err := c.Invoke("ModifyInstanceType", args, &response)
	if err == nil {
		return response.Return, nil
	}
	return false, err
}

type ModifyInstanceImageArgs struct {
	InstanceId       string
	ImageId          string
	InstancePassword string
}

type ModifyInstanceImageResponse struct {
	common.Response
	Return bool
}

// ModifyInstanceImage 更换或者重新安装实例操作系统
// You can read doc at https://docs.ksyun.com/read/latest/52/_book/oaModifyInstanceImage.html
func (c *Client) ModifyInstanceImage(args *ModifyInstanceImageArgs) (bool, error) {
	response := ModifyInstanceTypeResponse{}
	err := c.Invoke("ModifyInstanceImage", args, &response)
	if err == nil {
		return response.Return, nil
	}
	return false, err
}

type DescribeInstanceVncResponse struct {
	common.Response
	VncUrl string
}

// DescribeInstanceVnc 获取VNC信息
// You can read doc at https://docs.ksyun.com/read/latest/52/_book/oaDescribeInstanceVnc.html
func (c *Client) DescribeInstanceVnc(instanceId string) (string, error) {
	response := DescribeInstanceVncResponse{}
	err := c.Invoke("DescribeInstanceVnc", instanceId, &response)
	if err == nil {
		return response.VncUrl, nil
	}
	return "", err
}

type AttachNetworkInterfaceArgs struct {
	InstanceId       string
	SecurityGroupId  []string
	SubnetId         string
	PrivateIpAddress string
}

type AttachNetworkInterfaceResponse struct {
	common.Response
	Return bool
}

// AttachNetworkInterface 为主机添加网卡
// You can read doc at https://docs.ksyun.com/read/latest/52/_book/oaAttachNetworkInterface.html
func (c *Client) AttachNetworkInterface(args *AttachNetworkInterfaceArgs) (bool, error) {
	response := AttachNetworkInterfaceResponse{}
	err := c.Invoke("AttachNetworkInterface", args, &response)
	if err == nil {
		return response.Return, nil
	}
	return false, err
}

type ModifyNetworkInterfaceAttributeArgs struct {
	InstanceId         string
	NetworkInterfaceId string
	SecurityGroupId    []string
	SubnetId           string
	PrivateIpAddress   string
	DNS1               string
	DNS2               string
}

type ModifyNetworkInterfaceAttributeResponse struct {
	common.Response
	Return bool
}

// ModifyNetworkInterfaceAttribute 修改网络接口属性信息
// You can read doc at https://docs.ksyun.com/read/latest/52/_book/oaModifyNetworkInterfaceAttribute.html
func (c *Client) ModifyNetworkInterfaceAttribute(args *ModifyNetworkInterfaceAttributeArgs) (bool, error) {
	response := ModifyNetworkInterfaceAttributeResponse{}
	err := c.Invoke("ModifyNetworkInterfaceAttribute", args, &response)
	if err == nil {
		return response.Return, nil
	}
	return false, err
}

type DetachNetworkInterfaceArgs struct {
	InstanceId         string
	NetworkInterfaceId string
}
type DetachNetworkInterfaceResponse struct {
	common.Response
	Return bool
}

// DetachNetworkInterface 删除主机网络接口
// You can read doc at https://docs.ksyun.com/read/latest/52/_book/oaDetachNetworkInterface.html
func (c *Client) DetachNetworkInterface(args *DetachNetworkInterfaceArgs) (bool, error) {
	response := DetachNetworkInterfaceResponse{}
	err := c.Invoke("DetachNetworkInterface", args, &response)
	if err == nil {
		return response.Return, nil
	}
	return false, err
}
