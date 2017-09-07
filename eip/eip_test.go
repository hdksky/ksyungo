package eip

import (
	"fmt"
	"github.com/hdksky/ksyungo/slb"
	"testing"
)

var (
	testAccessKeyId     = "AKLTUxyeuc11TQ2gRi2yJN7FiA"
	testAccessKeySecret = "OHp/RfCLH+/c5rrH/+k0g9Mih3289ZonVyMITpordadELY6CzsECZcIc+X/oslbJCQ=="
)

func TestClient_AllocateAddress(t *testing.T) {
	client := NewClient(testAccessKeyId, testAccessKeySecret, regions[0])
	ls, err := client.GetLines()
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.AllocateAddress(&AllocateAddressArgs{
		LineId:     ls[0].LineId,
		RegionId:   regions[0],
		BandWidth:  1,
		ChargeType: ChargeType_Time,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_DescribeAddresses(t *testing.T) {
	client := NewClient(testAccessKeyId, testAccessKeySecret, regions[0])
	eips, err := client.DescribeAddresses(&DescribeAddressesArgs{})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", eips)
}

func TestClient_ReleaseAddress(t *testing.T) {
	client := NewClient(testAccessKeyId, testAccessKeySecret, regions[0])

	var eips []AddressDetail
	var err error
	if eips, err = client.DescribeAddresses(&DescribeAddressesArgs{}); err != nil {
		t.Fatal(err)

	}

	if len(eips) == 0 {
		TestClient_AllocateAddress(t)

		eips, err = client.DescribeAddresses(&DescribeAddressesArgs{})
		if err != nil {
			t.Fatal(err)
		}
	}

	if _, err := client.ReleaseAddress(&ReleaseAddressArgs{eips[0].AllocationId}); err != nil {
		t.Fatal(err)
	}
}

func TestClient_ModifyAddress(t *testing.T) {
	client := NewClient(testAccessKeyId, testAccessKeySecret, regions[0])

	var eips []AddressDetail
	var err error
	eips, err = client.DescribeAddresses(&DescribeAddressesArgs{})
	if err != nil {
		t.Fatal(err)
	}

	if len(eips) == 0 {
		TestClient_AllocateAddress(t)
	}

	eips, err = client.DescribeAddresses(&DescribeAddressesArgs{})
	if err != nil {
		t.Fatal(err)
	}

	var modifyArgs ModifyAddressArgs = ModifyAddressArgs{
		AllocationId: eips[0].AllocationId,
		BandWidth:    eips[0].BandWidth + 1,
	}

	if _, err := client.ModifyAddress(&modifyArgs); err != nil {
		t.Fatal(err)
	}
}

func TestClient_AssociateAddress(t *testing.T) {
	slbCli := slb.NewClient(testAccessKeyId, testAccessKeySecret, regions[0])
	slbs, err := slbCli.DescribeLoadBalancers(&slb.DescribeLoadBalancersArgs{
		State: slb.State_Disassociate,
	})
	if err != nil {
		t.Fatal(err)
	}

	eipCli := NewClient(testAccessKeyId, testAccessKeySecret, regions[0])
	eips, err := eipCli.DescribeAddresses(&DescribeAddressesArgs{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("-----slbs %v\n", slbs)
	fmt.Printf("-----eips %v\n", eips)
	//eipCli.AssociateAddress(&AssociateAddressArgs{})
}
