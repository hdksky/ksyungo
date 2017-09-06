package kec

import "testing"

func TestDescribeInstanceTypes(t *testing.T) {
	plans, err := DescribeInstanceTypes("cn-bejing-6a")
	if err != nil {
		t.Fatal(err)
	}

	if len(plans) == 0 {
		t.Fatal("describe instanceType")
	}
}
