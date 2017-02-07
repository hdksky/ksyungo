package eip

import (
	"time"

	"github.com/hdksky/ksyungo/common"
)

type AllocateAddressArgs struct {
	LineId       string
	BandWidth    int
	ChargeType   string
	PurchaseTime int
}

type AllocateAddressResponse struct {
	common.Response
	PublicIp     string
	AllocationId string
}

// AllocateAddress 创建EIP
// You can read doc at https://docs.ksyun.com/read/latest/57/_book/Action/AllocateAddress.html
func (c *Client) AllocateAddress(args *AllocateAddressArgs) (*AllocateAddressResponse, error) {
	response := AllocateAddressResponse{}
	err := c.Invoke("AllocateAddress", nil, &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}

type ReleaseAddressResponse struct {
	common.Response
	Return bool
}

// ReleaseAddress 创建EIP
// You can read doc at https://docs.ksyun.com/read/latest/57/_book/Action/ReleaseAddress.html
func (c *Client) ReleaseAddress(allocationId string) (bool, error) {
	response := ReleaseAddressResponse{}
	err := c.Invoke("ReleaseAddress", allocationId, &response)
	if err == nil {
		return response.Return, nil
	}
	return false, err
}

type AssociateAddressArgs struct {
	AllocationId       string
	InstanceType       string
	InstanceId         string
	NetworkInterfaceId string
}

type AssociateAddressResponse struct {
	common.Response
	Return bool
}

// AssociateAddress 绑定弹性IP
// You can read doc at https://docs.ksyun.com/read/latest/57/_book/Action/AssociateAddress.html
func (c *Client) AssociateAddress(args *AssociateAddressArgs) (bool, error) {
	response := AssociateAddressResponse{}
	err := c.Invoke("AssociateAddress", args, &response)
	return response.Return, err
}

type DisassociateAddressResponse struct {
	common.Response
	Return bool
}

// DisassociateAddress 解绑弹性IP
// You can read doc at https://docs.ksyun.com/read/latest/57/_book/Action/DisassociateAddress.html
func (c *Client) DisassociateAddress(allocationId string) (bool, error) {
	response := DisassociateAddressResponse{}
	err := c.Invoke("DisassociateAddress", allocationId, &response)
	return response.Return, err
}

type KV struct {
	Name  string
	Value string
}

type DescribeAddressesArgs struct {
	AllocationId []string
	Filter       []KV
	MaxResults   int
	NextToken    string
}

type AddressType struct {
	CreateTime         time.Time
	PublicIp           string
	AllocationId       string
	State              string
	LineId             string
	BandWidth          int
	InstanceType       string
	InstanceId         string
	NetworkInterfaceId string
	InternetGatewayId  string
	BandWidthShareId   string
	IsBandWidthShare   bool
}

type DescribeAddressesResponse struct {
	common.Response
	AddressesSet struct {
		Item []AddressType
	}
	NextToken string
}

// DescribeAddresses 描述EIP
// You can read doc at https://docs.ksyun.com/read/latest/57/_book/Action/DescribeAddresses.html
func (c *Client) DescribeAddresses(args *DescribeAddressesArgs) (*DescribeAddressesResponse, error) {
	response := DescribeAddressesResponse{}
	err := c.Invoke("DescribeAddresses", args, &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}
