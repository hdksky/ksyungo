package slb

import (
	"fmt"
	"testing"
)

func TestClient_DescribeInstancesWithListener(t *testing.T) {
	client := NewTestClientForDebug()
	bindings, err := client.DescribeInstancesWithListener(&DescribeInstancesWithListenerArgs{})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", bindings)
}

func TestClient_RegisterInstancesWithListener(t *testing.T) {
	client := NewTestClientForDebug()
	lns, err := client.DescribeListeners(&DescribeListenersArgs{})
	if err != nil {
		t.Fatal(err)
	}

	if len(lns) == 0 {
		t.Fatal("no listeners")
	}

	_, err = client.RegisterInstancesWithListener(&RegisterInstancesWithListenerArgs{
		ListenerId:     lns[0].ListenerId,
		RealServerIp:   "10.0.0.2",
		RealServerPort: 80,
		RealServerType: "host",
	})

	if err != nil {
		t.Fatal(err)
	}

}

func TestClient_DeregisterInstancesFromListener(t *testing.T) {
	client := NewTestClientForDebug()
	bindings, err := client.DescribeInstancesWithListener(&DescribeInstancesWithListenerArgs{})
	if err != nil {
		t.Fatal(err)
	}

	if len(bindings) == 0 {
		t.Fatal("register not found")
	}

	if ok, err := client.DeregisterInstancesFromListener(&DeregisterInstancesFromListenerArgs{
		RegisterId: bindings[0].RegisterId,
	}); err != nil {
		t.Fatal(err)
	} else if !ok {
		t.Fatal("DeregisterInstancesFromListener error")
	}
}
