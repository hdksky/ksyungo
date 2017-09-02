package eip

var (
	regions = []string{"cn-beijing-6", "cn-shanghai-2", "cn-hongkong-2"}

	ChargeType_PrePay string = "Monthly"
	ChargeType_Peak   string = "Peak"
	ChargeType_Time   string = "Daily"

	bandWidth_Min int = 1
	bandWidth_Max int = 200

	purchaseTime_Min int = 1
	purchaseTime_Max int = 36
)
