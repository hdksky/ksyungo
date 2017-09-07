package kec

import (
	"fmt"
	"testing"
)

func TestClient_DescribeImages(t *testing.T) {
	client := NewTestClientForDebug()
	imgs, err := client.DescribeImages("")
	if err != nil {
		t.Errorf("DescribeInstances failed: %v", err)
	}

	fmt.Printf("%#v\n", imgs)
}
