package kec

import "github.com/hdksky/aliyungo/common"

type InstanceMonitoringType struct {
	InstanceId string
	Monitoring struct {
		State string
	}
}

type MonitorInstancesArgs struct {
	InstanceId []string
}
type MonitorInstancesResponse struct {
	common.Response
	InstancesSet struct {
		Item []InstanceMonitoringType
	}
}

// MonitorInstances 启动实例监控
// You can read doc at https://docs.ksyun.com/read/latest/52/_book/oaMonitorInstances.html
func (c *Client) MonitorInstances(instanceId []string) ([]InstanceMonitoringType, error) {
	response := MonitorInstancesResponse{}
	args := &MonitorInstancesArgs{InstanceId: instanceId}
	err := c.Invoke("MonitorInstances", args, &response)
	if err == nil {
		return response.InstancesSet.Item, nil
	}
	return nil, err
}

type UnmonitorInstancesArgs struct {
	InstanceId []string
}
type UnmonitorInstancesResponse struct {
	common.Response
	InstancesSet struct {
		Item []InstanceMonitoringType
	}
}

// UnmonitorInstances 取消实例监控
// You can read doc at https://docs.ksyun.com/read/latest/52/_book/oaUnmonitorInstances.html
func (c *Client) UnmonitorInstances(instanceId []string) ([]InstanceMonitoringType, error) {
	response := UnmonitorInstancesResponse{}
	args := &UnmonitorInstancesArgs{InstanceId: instanceId}
	err := c.Invoke("UnmonitorInstances", args, &response)
	if err == nil {
		return response.InstancesSet.Item, nil
	}
	return nil, err
}
