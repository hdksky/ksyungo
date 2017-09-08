package region

import "errors"

var RegionToZonesInfo map[string][]Zone = map[string][]Zone{
	"cn-beijing-6": {
		{
			ZoneId:                        "cn-bejing-6a",
			LocalName:                     "可用区A",
			AvailableInstanceTypeFamilies: []string{"I1", "C1", "I2", "I1_E", "C1_E", "I2_E"},
		},
		{
			ZoneId:                        "cn-bejing-6b",
			LocalName:                     "可用区B",
			AvailableInstanceTypeFamilies: []string{"I2", "I2_E"},
		},
	},
	"cn-shanghai-2": {
		{
			ZoneId:                        "cn-shanghai-2a",
			LocalName:                     "可用区A",
			AvailableInstanceTypeFamilies: []string{"I1", "C1", "I2", "I1_E", "C1_E", "I2_E"},
		},
	},
	"cn-hongkong-2": {
		{
			ZoneId:                        "cn-hongkong-2a",
			LocalName:                     "可用区A",
			AvailableInstanceTypeFamilies: []string{"I1"},
		},
	},
}

func DescribeZones(region string) ([]Zone, error) {
	zones, ok := RegionToZonesInfo[region]
	if !ok {
		return nil, errors.New("region not found")
	}

	return zones, nil
}

type Zone struct {
	ZoneId                        string   `json:"ZoneId"`
	LocalName                     string   `json:"LocalName"`
	AvailableInstanceTypeFamilies []string `json:"AvailableInstanceTypeFamilies"`
}
