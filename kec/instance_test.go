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

	vpcCli := vpc.NewClient(TestAccessKeyId, TestAccessKeySecret, TestRegionId)
	if vpcs, err := vpcCli.DescribeVpcs(nil); err != nil {
		t.Fatal(err)
	} else if len(vpcs.VpcSet.Item) == 0 {
		t.Fatal("Vpcs not found")
	}

	sbs, err := vpcCli.DescribeSubnets(&vpc.DescribeSubnetsArgs{})
	if err != nil {
		t.Fatal(err)
	}

	if len(sbs) == 0 {
		t.Fatal("subnet not found")
	}

	sgs, err := vpcCli.DescribeSecurityGroups(&vpc.DescribeSecurityGroupsArgs{})
	if err != nil {
		t.Fatal(err)
	} else if len(sgs) == 0 {
		t.Fatal("SecurityGroups not found")
	}

	_, err = client.RunInstances(&RunInstancesArgs{
		ImageId:          "7aa79b22-a840-4836-a7ad-d440a0cf8bef",
		SecurityGroupId:  sgs[0].SecurityGroupId,
		SubnetId:         sbs[0].SubnetId,
		InstanceType:     "I1.1A",
		MinCount:         1,
		MaxCount:         1,
		InstancePassword: "Xiaowenrou123",
		InstanceName:     "michael",
		ChargeType:       "Monthly",
		DataDiskGb:       0,
		PurchaseTime:     1,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_DescribeInstances(t *testing.T) {
	client := NewTestClientForDebug()

	i, err := client.DescribeInstances(&DescribeInstancesArgs{})
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
