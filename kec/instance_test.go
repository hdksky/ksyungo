package kec

import (
	"fmt"
	"github.com/hdksky/ksyungo/vpc"
	"testing"
)

func TestClient_RunInstances(t *testing.T) {
	client := NewTestClientForDebug()

	images, err := client.DescribeImages("")
	if err != nil {
		t.Fatal(err)
	}

	if len(images) == 0 {
		t.Fatal("create instance image not found")
	}

	var vpcId string
	if vpcs, err := vpc.NewClient(TestAccessKeyId, TestAccessKeySecret, TestRegionId).
		DescribeVpcs(nil); err != nil {
		t.Fatal(err)
	} else if len(vpcs.VpcSet.Item) == 0 {
		t.Fatal("Vpcs not found")
	} else {
		vpcId = vpcs.VpcSet.Item[0].VpcId
	}

	var securityGroupId string
	if sgs, err := vpc.NewClient(TestAccessKeyId, TestAccessKeySecret, TestRegionId).
		DescribeSecurityGroups(&vpc.DescribeSecurityGroupsArgs{}); err != nil {
		t.Fatal(err)
	} else if len(sgs) == 0 {
		t.Fatal("SecurityGroups not found")
	} else {
		securityGroupId = sgs[0].VpcId
	}

	_, err = client.RunInstances(&RunInstancesArgs{
		ImageId:          images[0].ImageId,
		SecurityGroupId:  securityGroupId,
		SubnetId:         vpcId,
		MinCount:         1,
		MaxCount:         1,
		InstancePassword: "123456qweQWE",
		InstanceName:     "test",
		ChargeType:       "Daily",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_DescribeInstances(t *testing.T) {
	client := NewTestClientForDebug()

	i, err := client.DescribeInstances(&DescribeInstancesArgs{
		Filter: []KV{KV{Name: "instance-id", Value: []string{"123", "456"}}}})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", i)
}

func TestClient_AvailabilityZone(t *testing.T) {
	client := NewTestClientForDebug()

	if err := client.AvailabilityZone(&AvailabilityZoneArgs{
		Region: "cn-beijing-6",
	}); err != nil {
		t.Fatal(err)
	}
}
