/**
 * Author:  Nyxvectar Yan
 * Repo:    go-zju-formulas
 * Created: 07/23/2025
 */

package algebran

import (
	"errors"
	"math"
)

func valid(base, x float64) (bool, error) {
	if base <= 0 || base == 1 {
		var errNew = "对数的底须大于0且不为1"
		return false, errors.New(errNew)
	} else if x <= 0 {
		var errNew = "对数的真数须大于0"
		return false, errors.New(errNew)
	} else {
		return true, nil
	}
}

func LogBaseConvert(base, x float64) (float64, error) {
	var validOrNot, errNew = valid(base, x)
	if validOrNot != true {
		return 0, errNew
	} else {
		var (
			lnBase = math.Log(base)
			lnX    = math.Log(x)
		)
		return lnX / lnBase, nil
	}
}

func LogEquation(base, x float64) (float64, error) {
	var validOrNot, errNew = valid(base, x)
	if validOrNot != true {
		return 0, errNew
	} else {
		return x, nil
	}
}

func GrowthAvg(present, previous float64) (float64, error) {
	if present == 0 {
		var errNew = "起始值不得为零"
		return 0, errors.New(errNew)
	} else {
		var avG = (present - previous) / present
		return avG, nil
	}
}
