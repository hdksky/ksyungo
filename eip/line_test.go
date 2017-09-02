package eip

import (
	"testing"
)

func TestClient_GetAllLines(t *testing.T) {
	lines_res, err := GetAllLines(testAccessKeyId, testAccessKeySecret)
	if err != nil {
		t.Fatal(err)
	}

	if len(lines_res) != len(regions) {
		t.Fatal(lines_res)
	}

	lines_cached, err := GetAllLines(testAccessKeyId, testAccessKeySecret)
	if err != nil {
		t.Fatal(err)
	}

	if len(lines_cached) != len(lines_res) || len(lines_cached) == 0 {
		t.Fatal(lines_res, lines_cached)
	}
}

func TestClient_GetLines(t *testing.T) {
	client := NewClient(testAccessKeyId, testAccessKeySecret, "cn-hongkong-2")
	_, err := client.GetLines()
	if err != nil {
		t.Fatal(err)
	}
}
