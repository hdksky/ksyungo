package kec

import "errors"

//doc ref https://docs.ksyun.com/read/latest/52/_book/oashilitaocanleixing.html

var (
	zoneToPlansMapping map[string][]*Plan = map[string][]*Plan{
		"cn-bejing-6a":   []*Plan{Plan_I1, Plan_C1, Plan_I2, Plan_I1_Enhance, Plan_C1_Enhance, Plan_I2_Enhance},
		"cn-bejing-6b":   []*Plan{Plan_I2, Plan_I2_Enhance},
		"cn-shanghai-2a": []*Plan{Plan_I1, Plan_C1, Plan_I2, Plan_I1_Enhance, Plan_C1_Enhance, Plan_I2_Enhance},
		"cn-hongkong-2a": []*Plan{Plan_I1},
	}

	Plan_I1 *Plan = &Plan{
		Name:           "I1",
		ShowName:       "IO优化型I1",
		Definition:     "IO优化型I1实例是高磁盘IO的最佳选择，提供每秒数万次低延迟性随机 I/O 操作(IOPS)",
		UseCase:        "适合于低延时、I/O密集型应用。",
		SysVolumn:      "Linux操作系统 20GB、Windows 操作系统 50GB",
		Specifications: "收/发 10万PPS",
		Items: []PlanItem{
			{"I1.1A", 1, 1, 0, 50},
			{"I1.1B", 1, 2, 0, 50},
			{"I1.1C", 1, 4, 0, 50},
			{"I1.2A", 2, 2, 0, 200},
			{"I1.2B", 2, 4, 0, 200},
			{"I1.2C", 2, 8, 0, 200},
			{"I1.4A", 4, 4, 0, 500},
			{"I1.4B", 4, 8, 0, 500},
			{"I1.4C", 4, 16, 0, 500},
			{"I1.8A", 8, 8, 0, 800},
			{"I1.8B", 8, 16, 0, 800},
			{"I1.8C", 8, 32, 0, 800},
			{"I1.12A", 12, 12, 0, 1000},
			{"I1.12B", 12, 24, 0, 1000},
			{"I1.12C", 12, 48, 0, 1000},
			{"I1.16A", 16, 16, 0, 1200},
			{"I1.16B", 16, 32, 0, 1200},
			{"I1.16C", 16, 64, 0, 1200},
		},
	}

	Plan_C1 *Plan = &Plan{
		Name:           "C1",
		ShowName:       "计算优化型C1",
		Definition:     "采用Intel Xeon E5-2680 v3 (Haswell) 处理器，2.5GHz的主频，DDR4内存。是高磁盘IO的最佳选择，提供每秒数万次低延迟性随机I/O操作（IOPS）",
		UseCase:        "建议用于游戏服务器、数据库服务器。",
		SysVolumn:      "Linux操作系统 20GB、Windows 操作系统 50GB",
		Specifications: "收/发 10万PPS",
		Items: []PlanItem{
			{"C1.1A", 1, 1, 0, 50},
			{"C1.1B", 1, 2, 0, 50},
			{"C1.1C", 1, 4, 0, 50},
			{"C1.2A", 2, 2, 0, 200},
			{"C1.2B", 2, 4, 0, 200},
			{"C1.2C", 2, 8, 0, 200},
			{"C1.4A", 4, 4, 0, 500},
			{"C1.4B", 4, 8, 0, 500},
			{"C1.4C", 4, 16, 0, 500},
			{"C1.8A", 8, 8, 0, 800},
			{"C1.8B", 8, 16, 0, 800},
			{"C1.8C", 8, 32, 0, 800},
			{"C1.12A", 12, 12, 0, 1000},
			{"C1.12B", 12, 24, 0, 1000},
			{"C1.12C", 12, 48, 0, 1000},
			{"C1.16A", 16, 16, 0, 1200},
			{"C1.16B", 16, 32, 0, 1200},
			{"C1.16C", 16, 64, 0, 1200},
			{"C1.24A", 24, 24, 0, 1500},
			{"C1.24B", 24, 48, 0, 1500},
			{"C1.24C", 24, 32, 0, 1500},
			{"C1.24D", 24, 64, 0, 1500},
			{"C1.32A", 32, 32, 0, 2000},
			{"C1.32B", 32, 64, 0, 2000},
		},
	}

	Plan_I2 *Plan = &Plan{
		Name:           "I2",
		ShowName:       "IO优化型I2",
		Definition:     "采用Intel Xeon E5-2690 v4 (Broadwell)处理器，2.6GHz的高主频，DDR4内存。是高磁盘IO的最佳选择，提供每秒数万次低延迟性随机I/O操作（IOPS）",
		UseCase:        "适合于低延时、I/O密集型应用，建议用于游戏服务器、数据库服务器和高性能Web服务器等对I/O和计算性能均要求较高的场景。",
		SysVolumn:      "Linux操作系统 20GB、Windows 操作系统 50GB",
		Specifications: "收/发 10万PPS",
		Items: []PlanItem{
			{"I2.1A", 1, 1, 0, 50},
			{"I2.1B", 1, 2, 0, 50},
			{"I2.1C", 1, 4, 0, 50},
			{"I2.1D", 1, 8, 0, 50},
			{"I2.2B", 2, 4, 0, 200},
			{"I2.2C", 2, 8, 0, 200},
			{"I2.2D", 2, 16, 0, 200},
			{"I2.4B", 4, 8, 0, 500},
			{"I2.4C", 4, 16, 0, 500},
			{"I2.4D", 4, 32, 0, 500},
			{"I2.8B", 8, 16, 0, 800},
			{"I2.8C", 8, 32, 0, 800},
			{"I2.8D", 8, 64, 0, 800},
			{"I2.12B", 12, 24, 0, 1000},
			{"I2.12C", 12, 48, 0, 1000},
			{"I2.12D", 12, 96, 0, 1000},
			{"I2.16B", 16, 32, 0, 1200},
			{"I2.16C", 16, 64, 0, 1200},
			{"I2.16D", 16, 128, 0, 1200},
			{"I2.24B", 24, 48, 0, 1500},
			{"I2.24C", 24, 96, 0, 1500},
			{"I2.24D", 24, 192, 0, 1500},
			{"I2.32B", 32, 64, 0, 2000},
			{"I2.32C", 32, 128, 0, 2000},
			{"I2.32D", 32, 256, 1000, 3000},
			{"I2.48D", 48, 384, 2000, 4000},
		},
	}

	Plan_I1_Enhance *Plan = &Plan{
		Name:           "I1",
		ShowName:       "计算优化型I1联网增强",
		Definition:     "相比于前一代联网增强，联网增强2.0对基础架构做了优化调整。在精确管理PPS转发持久稳定的同时支持了更多云服务器类型和全类型操作系统。",
		UseCase:        "视频直播、即时通讯、房间式强联网网游等对网络实时性要求较高的应用。",
		SysVolumn:      "Linux操作系统 20GB、Windows 操作系统 50GB",
		Specifications: "收/发 30万PPS",
		Items: []PlanItem{
			{"I1.8A", 8, 8, 0, 800},
			{"I1.8B", 8, 16, 0, 800},
			{"I1.8C", 8, 32, 0, 800},
			{"I1.12A", 12, 12, 0, 1000},
			{"I1.12B", 12, 24, 0, 1000},
			{"I1.12C", 12, 48, 0, 1000},
			{"I1.16A", 16, 16, 0, 1200},
			{"I1.16B", 16, 32, 0, 1200},
			{"I1.16C", 16, 64, 0, 1200},
		},
	}

	Plan_C1_Enhance *Plan = &Plan{
		Name:           "C1",
		ShowName:       "IO优化型C1联网增强",
		Definition:     "相比于前一代联网增强，联网增强2.0对基础架构做了优化调整。在精确管理PPS转发持久稳定的同时支持了更多云服务器类型和全类型操作系统。",
		UseCase:        "视频直播、即时通讯、房间式强联网网游等对网络实时性要求较高的应用。",
		SysVolumn:      "Linux操作系统 20GB、Windows 操作系统 50GB",
		Specifications: "收/发 30万PPS",
		Items: []PlanItem{
			{"C1.8A", 8, 8, 0, 800},
			{"C1.8B", 8, 16, 0, 800},
			{"C1.8C", 8, 32, 0, 800},
			{"C1.12A", 12, 12, 0, 1000},
			{"C1.12B", 12, 24, 0, 1000},
			{"C1.12C", 12, 48, 0, 1000},
			{"C1.16A", 16, 16, 0, 1200},
			{"C1.16B", 16, 32, 0, 1200},
			{"C1.16C", 16, 64, 0, 1200},
			{"C1.24A", 24, 24, 0, 1500},
			{"C1.24B", 24, 48, 0, 1500},
			{"C1.24C", 24, 32, 0, 1500},
			{"C1.24D", 24, 64, 0, 1500},
			{"C1.32A", 32, 32, 0, 2000},
			{"C1.32B", 32, 64, 0, 2000},
		},
	}

	Plan_I2_Enhance *Plan = &Plan{
		Name:           "I2",
		ShowName:       "IO优化型I2联网增强",
		Definition:     "相比于前一代联网增强，联网增强2.0对基础架构做了优化调整。在精确管理PPS转发持久稳定的同时支持了更多云服务器类型和全类型操作系统。",
		UseCase:        "视频直播、即时通讯、房间式强联网网游等对网络实时性要求较高的应用。",
		SysVolumn:      "Linux操作系统 20GB、Windows 操作系统 50GB",
		Specifications: "收/发 30万PPS",
		Items: []PlanItem{
			{"I2.8B", 8, 16, 0, 800},
			{"I2.8C", 8, 32, 0, 800},
			{"I2.8D", 8, 64, 0, 800},
			{"I2.12B", 12, 24, 0, 1000},
			{"I2.12C", 12, 48, 0, 1000},
			{"I2.12D", 12, 96, 0, 1000},
			{"I2.16B", 16, 32, 0, 1200},
			{"I2.16C", 16, 64, 0, 1200},
			{"I2.16D", 16, 128, 0, 1200},
			{"I2.24B", 24, 48, 0, 1500},
			{"I2.24C", 24, 96, 0, 1500},
			{"I2.24D", 24, 192, 0, 1500},
			{"I2.32B", 32, 64, 0, 2000},
			{"I2.32C", 32, 128, 0, 2000},
			{"I2.32D", 32, 256, 1000, 3000},
			{"I2.48D", 48, 384, 2000, 4000},
		},
	}
)

type PlanItem struct {
	Name      string
	Cpu       uint8
	Memory    uint16
	VolumnMin uint16
	VolumnMax uint16
}

type Plan struct {
	Name           string
	ShowName       string
	Definition     string
	UseCase        string
	Specifications string
	SysVolumn      string
	Items          []PlanItem
}

//ref https://docs.ksyun.com/read/latest/52/_book/oashilitaocanleixing.html
func DescribeInstanceTypes(zone string) ([]*Plan, error) {
	plans, ok := zoneToPlansMapping[zone]
	if !ok {
		return nil, errors.New("zone not exist")
	}
	return plans, nil
}
