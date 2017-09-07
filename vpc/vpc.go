package vpc

import "github.com/hdksky/ksyungo/common"

type VpcType struct {
	CreateTime string
	VpcName    string
	VpcId      string
	CidrBlock  string
	IsDefault  bool
}

type DescribeVpcsArgs struct {
	VpcIds []string
}

type DescribeVpcsResponse struct {
	common.Response
	VpcSet struct {
		Item []VpcType `xml:"item"`
	}
}

// DescribeVpcs describe vpc info
// You can read doc at https://docs.ksyun.com/read/latest/56/_book/Action/Vpcs/DescribeVpcs.html
func (c *Client) DescribeVpcs(vpcIds []string) (*DescribeVpcsResponse, error) {
	response := DescribeVpcsResponse{}
	args := &DescribeVpcsArgs{VpcIds: vpcIds}
	err := c.Invoke("DescribeVpcs", args, &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}

type Kv struct {
	Name  string
	Value []string
}

type DescribeSubnetsArgs struct {
	SubnetId []string `opt:"N"`
	Filter   []Kv
}

type SubnetType struct {
	CreateTime   string
	VpcId        string
	SubnetId     string
	SubnetType   string
	SubnetName   string
	CidrBlock    string
	DhcpIpFrom   string
	DhcpIpTo     string
	GatewayIp    string
	Dns1         string
	Dns2         string
	NetworkAclId string
	NatId        string
}

type DescribeSubnetsResonse struct {
	common.Response
	SubnetSet struct {
		Item []SubnetType `xml:"item"`
	}
}

// DescribeSubnets describe subnets info
// You can read doc at https://docs.ksyun.com/read/latest/56/_book/Action/Subnets/DescribeSubnets.html
func (c *Client) DescribeSubnets(args *DescribeSubnetsArgs) ([]SubnetType, error) {
	response := DescribeSubnetsResonse{}
	err := c.Invoke("DescribeSubnets", args, &response)
	if err != nil {
		return nil, err
	}

	return response.SubnetSet.Item, nil
}

type DescribeSecurityGroupsArgs struct {
	SecurityGroupId []string
	Filter          []Kv
}

type SecurityGroup struct {
	CreateTime            string
	VpcId                 string
	SecurityGroupName     string
	SecurityGroupId       string
	SecurityGroupType     string
	SecurityGroupEntrySet struct {
		Item []EntrySet `xml:"item"`
	}
}

type EntrySet struct {
	CidrBlock            string
	Direction            string
	PortRangeTo          uint16
	PortRangeFrom        uint16
	Protocol             string
	SecurityGroupEntryId string
}

type DescribeSecurityGroupsResponse struct {
	common.Response
	SecurityGroupSet struct {
		Item []SecurityGroup `xml:"item"`
	}
}

// DescribeSecurityGroups describe security groups
// You can read doc at https://docs.ksyun.com/read/latest/56/_book/Action/SecurityGroups/DescribeSecurityGroups.html
func (c *Client) DescribeSecurityGroups(args *DescribeSecurityGroupsArgs) ([]SecurityGroup, error) {
	response := DescribeSecurityGroupsResponse{}
	if err := c.Invoke("DescribeSecurityGroups", args, &response); err != nil {
		return nil, err
	}

	return response.SecurityGroupSet.Item, nil
}
