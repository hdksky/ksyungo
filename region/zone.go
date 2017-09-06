package region

import "errors"

type Zone struct {
	ZoneId   string `json:"ZoneId"`
	ZoneName string `json:"ZoneName"`
}

var Zones map[string][]Zone = map[string][]Zone{
	"cn-beijing-6": []Zone{
		{ZoneId: "cn-bejing-6a", ZoneName: "可用区A"},
		{ZoneId: "cn-bejing-6b", ZoneName: "可用区B"},
	},
	"cn-shanghai-2": []Zone{
		{ZoneId: "cn-shanghai-2a", ZoneName: "可用区A"},
	},
	"cn-hongkong-2": []Zone{
		{ZoneId: "cn-hongkong-2a", ZoneName: "可用区A"},
	},
}

func DescribeZones(region string) ([]Zone, error) {
	zones, ok := Zones[region]
	if !ok {
		return nil, errors.New("region not found")
	}

	return zones, nil
}
