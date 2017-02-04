package slb

import "testing"

func TestDescribeLoadBalancer(t *testing.T) {
	client := NewTestClientForDebug()
	lbs, err := client.DescribeLoadBalancers(&DescribeLoadBalancersArgs{})
	if err != nil {
		t.Errorf("DescribeLoadBalancers failed: %v", err)
		return
	}
	t.Logf("lbs %+v\n", lbs)
}
