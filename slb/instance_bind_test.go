package slb

import (
	"fmt"
	"github.com/hdksky/ksyungo/kec"
	"testing"
)

func TestClient_RegisterInstancesWithListener(t *testing.T) {
	sbCli := NewTestClientForDebug()
	kecCli := kec.NewClient(TestAccessKeyId, TestAccessKeySecret, TestRegion)

start:
	var lns []ListenerType
	var err error

	var instanceFound, lnsFound bool
	var loadBalancerId string

	//查询所有负载均衡列表
	lbs, err := sbCli.DescribeLoadBalancers(&DescribeLoadBalancersArgs{})
	if err != nil {
		t.Fatal(err)
	} else if len(lbs) == 0 {
		// TODO create LoadBalancer
	}

	for _, lb := range lbs {
		targetLbId, targetVpcId := lb.LoadBalancerId, lb.VpcId

		//根据负载均衡的ｖｐｃId查询在相同ＶＰＣ的主机
		if vpcInstances, err := kecCli.DescribeInstances(&kec.DescribeInstancesArgs{
			Filter: []kec.KV{
				{
					Name:  "vpc-id",
					Value: []string{targetVpcId},
				},
			},
		}); err != nil {
			continue
		} else if len(vpcInstances.InstancesSet.Item) == 0 {
			continue
		} else {
			instanceFound = true
			loadBalancerId = targetLbId

			//查询当前负载均衡的监听器列表
			if lns, err = sbCli.DescribeListeners(&DescribeListenersArgs{
				Filter: []KV{{Name: "load-balancer-id", Value: []string{targetLbId}}},
			}); err != nil {
				t.Fatal(err)
			} else if len(lns) == 0 {
				continue
			} else {
				lnsFound = true
				sbCli.RegisterInstancesWithListener(&RegisterInstancesWithListenerArgs{
					ListenerId:     lns[0].ListenerId,
					RealServerType: "host",
					RealServerIp:   vpcInstances.InstancesSet.Item[0].PrivateIpAddress,
					RealServerPort: 8080,
					Weight:         16,
				})
			}
		}
	}

	if !lnsFound {
		if _, err = sbCli.CreateListeners(&CreateListenersArgs{
			LoadBalancerId:           loadBalancerId,
			ListenerState:            "start",
			ListenerProtocol:         "TCP",
			ListenerPort:             80,
			Method:                   ListenerMethod_RoundRobin,
			SessionState:             "start",
			SessionPersistencePeriod: 3000,
			CookieType:               CookieType_ImplantCookie,
		}); err != nil {
			t.Fatal(err)
		}
		goto start
	}

	if !instanceFound {
		t.Fatal("instance not found")
	}
}

func TestClient_DescribeInstancesWithListener(t *testing.T) {
	client := NewTestClientForDebug()
	bindings, err := client.DescribeInstancesWithListener(&DescribeInstancesWithListenerArgs{})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", bindings)
}

func TestClient_RegisterInstancesWithListener_Simple(t *testing.T) {
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

func TestClient_DeregisterInstancesFromListener_Simple(t *testing.T) {
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
