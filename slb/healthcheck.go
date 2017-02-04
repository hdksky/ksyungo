package slb

import "github.com/hdksky/ksyungo/common"

type HealthCheckType struct {
	HealthCheckId      string
	HealthCheckState   string
	HealthyThreshold   int
	Interval           int
	Timeout            int
	UnhealthyThreshold int
}

type ConfigureHealthCheckArgs struct {
	ListenerId         string
	HealthCheckState   string
	HealthyThreshold   int
	Interval           int
	Timeout            int
	UnhealthyThreshold int
	UrlPath            string
	IsDefaultHostName  bool
	HostName           string
}

type ConfigureHealthCheckResponse struct {
	common.Response
	HealthCheckType
}

// ConfigureHealthCheck Configure health checker
// You can read doc at https://docs.ksyun.com/read/latest/55/_book/Action/ConfigureHealthCheck.html
func (c *Client) ConfigureHealthCheck(args *ConfigureHealthCheckArgs) (*ConfigureHealthCheckResponse, error) {
	response := ConfigureHealthCheckResponse{}
	err := c.Invoke("ConfigureHealthCheck", args, &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}

type ModifyHealthCheckArgs struct {
	HealthCheckId      string
	HealthCheckState   string
	HealthyThreshold   string
	Interval           string
	Timeout            int
	UnhealthyThreshold int
	UrlPath            string
	IsDefaultHostName  bool
	HostName           string
}

type ModifyHealthCheckResponse struct {
	common.Response
	HealthCheckType
}

// ModifyHealthCheck modify health checker
// You can read doc at https://docs.ksyun.com/read/latest/55/_book/Action/ModifyHealthCheck.html
func (c *Client) ModifyHealthCheck(args *ModifyHealthCheckArgs) (*ModifyHealthCheckResponse, error) {
	response := ModifyHealthCheckResponse{}
	err := c.Invoke("ModifyHealthCheck", args, &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}

type DeleteHealthCheckResponse struct {
	common.Response
	Return bool
}

// DeleteHealthCheck modify health checker
// You can read doc at https://docs.ksyun.com/read/latest/55/_book/Action/DeleteHealthCheck.html
func (c *Client) DeleteHealthCheck(healthCheckId string) (bool, error) {
	response := DeleteHealthCheckResponse{}
	err := c.Invoke("DeleteHealthCheck", healthCheckId, &response)
	if err == nil {
		return response.Return, nil
	}
	return false, err
}

type DescribeHealthChecksArgs struct {
	HealthCheckId []string
	Filter        []KV
}

type DescribeHealthChecksResponse struct {
	common.Response
	HealthCheckSet struct {
		Item []HealthCheckType
	}
}

// DescribeHealthChecks modify health checker
// You can read doc at https://docs.ksyun.com/read/latest/55/_book/Action/DescribeHealthChecks.html
func (c *Client) DescribeHealthChecks(args *DescribeHealthChecksArgs) ([]HealthCheckType, error) {
	response := DescribeHealthChecksResponse{}
	err := c.Invoke("DescribeHealthChecks", args, &response)
	if err == nil {
		return response.HealthCheckSet.Item, nil
	}
	return nil, err
}
