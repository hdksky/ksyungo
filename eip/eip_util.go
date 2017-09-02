package eip

import (
	"github.com/hdksky/ksyungo/common"
	"github.com/hdksky/ksyungo/util"
)

type AllocateAddressArgs struct {
	RegionId     string `json:"regionId"`
	LineId       string `json:"lineId"`
	BandWidth    int    `json:"bandWidth"`
	ChargeType   string `json:"chargeType"`
	PurchaseTime int    `json:"purchaseTime"` //unit month
}

func (a AllocateAddressArgs) validate() error {
	if len(a.RegionId) == 0 {
		return util.ParamNotFoundErr("regionId")
	}

	if !validateRegion(a.RegionId) {
		return util.ParamInvalid("regionId")
	}

	if len(a.LineId) == 0 {
		return util.ParamNotFoundErr("lineId")
	}

	if len(a.ChargeType) == 0 {
		return util.ParamNotFoundErr("chargeType")
	}

	if !validateChargeType(a.ChargeType) {
		return util.ParamInvalid("chargeType")
	}

	if a.BandWidth < bandWidth_Min || a.BandWidth > bandWidth_Max {
		return util.ParamValueInvalid("bandWidth", bandWidth_Min, bandWidth_Max)
	}

	if a.ChargeType == ChargeType_PrePay && (a.PurchaseTime < purchaseTime_Min || a.PurchaseTime > purchaseTime_Max) {
		return util.ParamValueInvalid("purchaseTime", purchaseTime_Min, purchaseTime_Max)
	}

	return nil
}

func (a ModifyAddressArgs) validate() error {
	if len(a.AllocationId) == 0 {
		return util.ParamNotFoundErr("allocationId")
	}

	if a.BandWidth < bandWidth_Min || a.BandWidth > bandWidth_Max {
		return util.ParamValueInvalid("bandWidth", bandWidth_Min, bandWidth_Max)
	}

	return nil
}

type AllocateAddressResponse struct {
	common.Response
	PublicIp     string
	AllocationId string
}

func validateRegion(region string) (ok bool) {
	for _, rg := range regions {
		if rg == region {
			ok = true
			return
		}
	}

	return
}

func validateChargeType(ct string) (ok bool) {
	return ct == ChargeType_Peak || ct == ChargeType_Time || ct == ChargeType_PrePay
}
