package kec

import "testing"

func TestClient_DescribeImages(t *testing.T) {
	client := NewTestClientForDebug()
	_, err := client.DescribeImages("")
	if err != nil {
		t.Errorf("DescribeInstances failed: %v", err)
	}
}
