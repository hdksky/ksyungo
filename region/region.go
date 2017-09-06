package region

// Region Ksyun region
type Region struct {
	ID   string
	Name string
}

// Regions 金山云支持的机房
var Regions = []Region{
	{"cn-beijing-6", "北京6区"},
	{"cn-shanghai-2", "上海2区"},
	{"cn-hongkong-2", "香港2区"},
}

func DescribeRegions() []Region {
	return Regions
}
