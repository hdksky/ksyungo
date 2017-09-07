package vpc

import (
	"fmt"
	"testing"
)

func TestClient_DescribeVpcs(t *testing.T) {
	cli := NewClient("AKLTUxyeuc11TQ2gRi2yJN7FiA", "OHp/RfCLH+/c5rrH/+k0g9Mih3289ZonVyMITpordadELY6CzsECZcIc+X/oslbJCQ==", "cn-beijing-6")
	cli.SetDebug(true)
	vpcs, err := cli.DescribeVpcs(nil)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", vpcs)
}

func TestClient_DescribeSubnets(t *testing.T) {
	cli := NewClient("AKLTUxyeuc11TQ2gRi2yJN7FiA", "OHp/RfCLH+/c5rrH/+k0g9Mih3289ZonVyMITpordadELY6CzsECZcIc+X/oslbJCQ==", "cn-beijing-6")
	vpcs, err := cli.DescribeVpcs(nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(vpcs.VpcSet.Item) == 0 {
		t.Fatal("subnets vpc not found")
	}

	var vpcIds []string
	for _, vpc := range vpcs.VpcSet.Item {
		vpcIds = append(vpcIds, vpc.VpcId)
	}

	if _, err := cli.DescribeSubnets(&DescribeSubnetsArgs{
		Filter: []Kv{{Name: "vpc-id", Value: vpcIds}},
	}); err != nil {
		t.Fatal(err)
	}
}

func TestClient_DescribeSecurityGroups(t *testing.T) {
	cli := NewClient("AKLTUxyeuc11TQ2gRi2yJN7FiA", "OHp/RfCLH+/c5rrH/+k0g9Mih3289ZonVyMITpordadELY6CzsECZcIc+X/oslbJCQ==", "cn-beijing-6")
	if _, err := cli.DescribeSecurityGroups(&DescribeSecurityGroupsArgs{}); err != nil {
		t.Fatal(err)
	}
}
