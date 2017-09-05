package kec

import (
	"fmt"
	"github.com/hdksky/ksyungo/util"
	"regexp"
)

var (
	uuidRegexp     string = `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`
	passWordRegexp string = `^[0-9a-zA-z]{8,32}$`
)

func (a RunInstancesArgs) validate() error {
	if len(a.ImageId) == 0 {
		return util.ParamNotFoundErr("ImageId")
	}

	if _, err := regexp.MatchString(uuidRegexp, a.ImageId); err != nil {
		return util.ParamInvalid("ImageId")
	}

	if a.MaxCount <= 0 || a.MaxCount < a.MinCount {
		return util.ParamInvalid("MaxCount")
	}

	if a.MinCount <= 0 || a.MaxCount < a.MinCount {
		return util.ParamInvalid("MinCount	")
	}

	if len(a.SubnetId) == 0 {
		return util.ParamNotFoundErr("SubnetId")
	}

	if _, err := regexp.MatchString(uuidRegexp, a.SubnetId); err != nil {
		return util.ParamInvalid("SubnetId")
	}

	if len(a.InstancePassword) == 0 {
		return util.ParamNotFoundErr("InstancePassword")
	}

	//if len(a.InstancePassword) < 8 || len(a.InstancePassword) > 32 {
	//	return util.ParamValueInvalid("InstancePassword", 8, 32)
	//}

	if _, err := regexp.MatchString(passWordRegexp, a.InstancePassword); err != nil {
		fmt.Println(err)
		return util.ParamInvalid("InstancePassword")
	}

	if len(a.ChargeType) == 0 {
		return util.ParamInvalid("ChargeType")
	}

	if a.ChargeType != "Monthly" && a.ChargeType != "Daily" {
		return util.ParamInvalid("ChargeType")
	}

	if a.ChargeType == "Monthly" && (a.PurchaseTime < 1 || a.PurchaseTime > 36) {
		return util.ParamValueInvalid("PurchaseTime", 1, 36)
	}

	if len(a.InstanceName) != 0 && (len(a.InstanceName) < 2 || len(a.InstanceName) > 64) {
		return util.ParamValueInvalid("InstanceName", 2, 64)
	}

	return nil
}
