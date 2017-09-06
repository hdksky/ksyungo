package region

import "testing"

func TestDescribeZones(t *testing.T) {
	zones, err := DescribeZones("cn-beijing-6")
	if err != nil {
		t.Fatal(err)
	}

	if len(zones) != 2 {
		t.Fatal(zones)
	}
}
