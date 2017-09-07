package slb

import (
	"github.com/hdksky/ksyungo/region"
	"github.com/hdksky/ksyungo/vpc"
	"testing"
)

func TestClient_CreateLoadBalancer_Pub(t *testing.T) {
	vpcCli := vpc.NewClient(TestAccessKeyId, TestAccessKeySecret, region.DescribeRegions()[0].ID)
	vpcs, err := vpcCli.DescribeVpcs(nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(vpcs.VpcSet.Item) == 0 {
		t.Fatal("creater load balaner err")
	}

	slbCli := NewTestClient()
	slbCli.SetDebug(true)

	lbs, err := slbCli.CreateLoadBalancer(&CreateLoadBalancerArgs{
		VpcId:            vpcs.VpcSet.Item[0].VpcId,
		LoadBalancerName: "apiCreateName",
		Type:             LoadBalancerType_Public,
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(lbs) == 0 {
		t.Fatal("creater load balaner err")
	}
}

func TestClient_CreateLoadBalancer_Internal(t *testing.T) {
	vpcCli := vpc.NewClient(TestAccessKeyId, TestAccessKeySecret, region.DescribeRegions()[0].ID)
	vpcs, err := vpcCli.DescribeVpcs(nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(vpcs.VpcSet.Item) == 0 {
		t.Fatal("creater load balaner err, vpc not find")
	}

	sns, err := vpcCli.DescribeSubnets(&vpc.DescribeSubnetsArgs{
		Filter: []vpc.Kv{vpc.Kv{Name: "vpc-id", Value: []string{vpcs.VpcSet.Item[0].VpcId}}}})
	if err != nil {
		t.Fatal(err)
	}

	if len(sns) == 0 {
		t.Fatal("creater load balaner err, subnets not find")
	}

	slbCli := NewTestClient()
	slbCli.SetDebug(true)
	lbs, err := slbCli.CreateLoadBalancer(&CreateLoadBalancerArgs{
		VpcId:            vpcs.VpcSet.Item[0].VpcId,
		LoadBalancerName: "apiCreateName",
		Type:             LoadBalancerType_Internal,
		SubnetId:         sns[0].SubnetId,
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(lbs) == 0 {
		t.Fatal("creater load balaner err")
	}
}

func TestClient_DeleteLoadBalancer(t *testing.T) {
	client := NewTestClient()

	lbs, err := client.DescribeLoadBalancers(&DescribeLoadBalancersArgs{
		State: State_Disassociate,
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(lbs) == 0 {
		t.Fatal("delete loadBalancer not found")
	}

	ok, err := client.DeleteLoadBalancer(&DeleteLoadBalancerArgs{
		LoadBalancerId: lbs[0].LoadBalancerId})
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("delete loadBalancer error")
	}
}

func TestClient_DescribeLoadBalancers(t *testing.T) {
	client := NewTestClientForDebug()

	var allLbs, assLbs, disassLbs []LoadBalancerDescription
	var err error

	if allLbs, err = client.DescribeLoadBalancers(&DescribeLoadBalancersArgs{}); err != nil {
		t.Fatal(err)
	}

	if assLbs, err = client.DescribeLoadBalancers(&DescribeLoadBalancersArgs{
		State: State_Associate,
	}); err != nil {
		t.Fatal(err)
	}

	if disassLbs, err = client.DescribeLoadBalancers(&DescribeLoadBalancersArgs{
		State: State_Disassociate,
	}); err != nil {
		t.Fatal(err)
	}

	if len(allLbs) != len(assLbs)+len(disassLbs) {
		t.Fatal("describe slb")
	}
}

func TestClient_ModifyLoadBalancer_Name(t *testing.T) {
	client := NewTestClientForDebug()

	var lbs []LoadBalancerDescription
	var err error
	if lbs, err = client.DescribeLoadBalancers(&DescribeLoadBalancersArgs{
		State: State_Disassociate,
	}); err != nil {
		t.Fatal(err)
	}

	if len(lbs) == 0 {
		t.Fatal("modify loadBalancer not found")
	}

	rps, err := client.ModifyLoadBalancer(&ModifyLoadBalancerArgs{
		LoadBalancerId:   lbs[0].LoadBalancerId,
		LoadBalancerName: "123",
	})
	if err != nil {
		t.Fatal(err)
	}

	if rps.LoadBalancerName != "123" {
		t.Fatal("modify loadBalancer　Name err")
	}
}

func TestClient_ModifyLoadBalancer_State(t *testing.T) {
	client := NewTestClientForDebug()

	var lbs []LoadBalancerDescription
	var err error
	if lbs, err = client.DescribeLoadBalancers(&DescribeLoadBalancersArgs{
		State: State_Associate,
	}); err != nil {
		t.Fatal(err)
	}

	if len(lbs) == 0 {
		t.Fatal("modify loadBalancer not found")
	}

	var state string
	if lbs[0].LoadBalancerState == LoadBalancerState_Start {
		state = LoadBalancerState_Stop
	} else {
		state = LoadBalancerState_Start
	}
	rps, err := client.ModifyLoadBalancer(&ModifyLoadBalancerArgs{
		LoadBalancerId:    lbs[0].LoadBalancerId,
		LoadBalancerState: state,
	})
	if err != nil {
		t.Fatal(err)
	}

	if rps.LoadBalancerState != state {
		t.Fatal("modify loadBalancer　Name err")
	}
}
