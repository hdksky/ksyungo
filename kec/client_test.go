package kec

import (
	"testing"
)

func TestGenerateClientToken(t *testing.T) {
	client := NewTestClient()
	for i := 0; i < 10; i++ {
		t.Log("GenerateClientToken: ", client.GenerateClientToken())
	}

}

func TestKECDescribe(t *testing.T) {

	client := NewTestClientForDebug()

	hosts, err := client.DescribeInstances(&DescribeInstancesArgs{})
	if err != nil {
		t.Errorf("DescribeInstances failed: %v", err)
	}

	t.Logf("Hosts: %+v", hosts)

}
