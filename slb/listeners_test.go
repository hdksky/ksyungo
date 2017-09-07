package slb

import (
	"testing"
)

func TestDescribeLoadBalancer(t *testing.T) {
	client := NewTestClientForDebug()
	lbs, err := client.DescribeLoadBalancers(&DescribeLoadBalancersArgs{})
	if err != nil {
		t.Errorf("DescribeLoadBalancers failed: %v", err)
		return
	}
	t.Logf("lbs %+v\n", lbs)
}

func TestClient_CreateListeners_TCP(t *testing.T) {
	client := NewTestClientForDebug()
	client.SetDebug(true)
	lbs, err := client.DescribeLoadBalancers(&DescribeLoadBalancersArgs{})
	if err != nil {
		t.Errorf("DescribeLoadBalancers failed: %v", err)
		return
	}

	if len(lbs) == 0 {
		t.Fatal("create loadBalancer not found")
	}

	_, err = client.CreateListeners(&CreateListenersArgs{
		LoadBalancerId:           lbs[0].LoadBalancerId,
		ListenerState:            "start",
		ListenerProtocol:         "TCP",
		ListenerPort:             80,
		Method:                   ListenerMethod_RoundRobin,
		SessionState:             "start",
		SessionPersistencePeriod: 3000,
		CookieType:               CookieType_ImplantCookie,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_CreateListeners_HTTP(t *testing.T) {
	client := NewTestClientForDebug()
	//client.SetDebug(true)
	lbs, err := client.DescribeLoadBalancers(&DescribeLoadBalancersArgs{})
	if err != nil {
		t.Errorf("DescribeLoadBalancers failed: %v", err)
		return
	}

	if len(lbs) == 0 {
		t.Fatal("create loadBalancer not found")
	}

	for _, lb := range lbs {
		if lb.Type == LoadBalancerType_Public {
			_, err = client.CreateListeners(&CreateListenersArgs{
				LoadBalancerId:           lb.LoadBalancerId,
				ListenerState:            "start",
				ListenerProtocol:         "HTTP",
				ListenerPort:             8003,
				Method:                   ListenerMethod_RoundRobin,
				SessionState:             "start",
				SessionPersistencePeriod: 3000,
				CookieType:               CookieType_ImplantCookie,
			})
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}

func TestClient_DeleteListeners(t *testing.T) {
	client := NewTestClientForDebug()
	//client.SetDebug(true)
	lns, err := client.DescribeListeners(&DescribeListenersArgs{})
	if err != nil {
		t.Errorf("describe listeners failed: %v", err)
		return
	}

	if len(lns) == 0 {
		t.Fatal("delete listeners not found")
	}

	ok, err := client.DeleteListeners(&DeleteListenerArgs{
		ListenerId: lns[0].ListenerId,
	})

	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Fatal("delete listener failed")
	}
}

func TestClient_ModifyListeners(t *testing.T) {
	client := NewTestClientForDebug()
	lns, err := client.DescribeListeners(&DescribeListenersArgs{})
	if err != nil {
		t.Errorf("describe listeners failed: %v", err)
		return
	}

	if len(lns) == 0 {
		t.Fatal("delete listeners not found")
	}

	target, state := lns[0], lns[0].ListenerState
	if state == ListenerState_Start {
		state = ListenerState_Stop
	} else {
		state = ListenerState_Start
	}

	if _, err := client.ModifyListeners(&ModifyListenersArgs{
		ListenerId:    target.ListenerId,
		ListenerState: state,
	}); err != nil {
		t.Fatal(err)
	}
}
