package region

import "testing"

func TestDescribeRegions(t *testing.T) {
	rs := DescribeRegions()
	if len(rs) == 0 {
		t.Fatal("describe regions")
	}
}
