package eip

import (
	"fmt"
	"os"

	"github.com/hdksky/ksyungo/common"
)

// DefaultWaitForInterval Interval for checking status in WaitForXXX method
const DefaultWaitForInterval = 5

// DefaultTimeout Default timeout value for WaitForXXX method
const DefaultTimeout = 60

const service = "eip"

type Client struct {
	common.Client
}

const (
	DefaultEndpoint = "https://eip.%s.api.ksyun.com"
	APIVersion      = "2016-03-04"
)

// NewClient creates a new instance of EIP client
func NewClient(accessKeyId, accessKeySecret, region string) *Client {
	endpoint := os.Getenv("EIP_ENDPOINT")
	if endpoint == "" {
		endpoint = fmt.Sprintf(DefaultEndpoint, region)
	}
	return NewClientWithEndpoint(endpoint, accessKeyId, accessKeySecret, region, service)
}

func NewClientWithEndpoint(endpoint string, accessKeyId, accessKeySecret, region, service string) *Client {
	client := &Client{}
	client.Init(endpoint, APIVersion, accessKeyId, accessKeySecret, region, service)
	return client
}
