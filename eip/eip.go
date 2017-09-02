package eip

import (
	"errors"
	"github.com/hdksky/ksyungo/common"
)

// AllocateAddress 创建EIP
// You can read doc at https://docs.ksyun.com/read/latest/57/_book/Action/AllocateAddress.html
func (c *Client) AllocateAddress(args *AllocateAddressArgs) (*AllocateAddressResponse, error) {
	if args == nil {
		return nil, errors.New("create eip param not found")
	}
	if err := args.validate(); err != nil {
		return nil, err
	}

	response := &AllocateAddressResponse{}
	err := c.Invoke("AllocateAddress", args, response)
	if err == nil {
		return response, nil
	}
	return nil, err
}

type ReleaseAddressArgs struct {
	AllocationId string
}

type ReleaseAddressResponse struct {
	common.Response
	Return bool
}

// ReleaseAddress 创建EIP
// You can read doc at https://docs.ksyun.com/read/latest/57/_book/Action/ReleaseAddress.html
func (c *Client) ReleaseAddress(args *ReleaseAddressArgs) (bool, error) {
	response := ReleaseAddressResponse{}
	err := c.Invoke("ReleaseAddress", args, &response)
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

type AddressDetail struct {
	CreateTime         string
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
		Item []AddressDetail `xml:"item"`
	}
	NextToken string
}

// DescribeAddresses 描述EIP
// You can read doc at https://docs.ksyun.com/read/latest/57/_book/Action/DescribeAddresses.html
func (c *Client) DescribeAddresses(args *DescribeAddressesArgs) ([]AddressDetail, error) {
	var response DescribeAddressesResponse
	err := c.Invoke("DescribeAddresses", args, &response)

	return response.AddressesSet.Item, err
}

type ModifyAddressArgs struct {
	AllocationId string `json:"allocationId"`
	BandWidth    int    `json:"bandWidth"`
}

type ModifyAddressResponse AddressDetail

// ModifyAddress 更新EIP
// doc ref https://docs.ksyun.com/read/latest/57/_book/Action/ModifyAddress.html
func (c *Client) ModifyAddress(args *ModifyAddressArgs) (*AddressDetail, error) {
	var response ModifyAddressResponse
	if err := args.validate(); err != nil {
		return nil, err
	}

	if err := c.Invoke("ModifyAddress", args, &response); err != nil {
		return nil, err
	}

	newAddr := AddressDetail(response)
	return &newAddr, nil
}
