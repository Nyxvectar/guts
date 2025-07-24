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
		return false, errors.New("对数的底须大于0且不为1")
	}
	if x <= 0 {
		return false, errors.New("对数的真数须大于0")
	}
	return true, nil
}

func LogBaseConvert(base, x float64) (float64, error) {
	if ok, err := valid(base, x); !ok {
		return 0, err
	}
	return math.Log(x) / math.Log(base), nil
}

func LogEquation(base, x float64) (float64, error) {
	if ok, err := valid(base, x); !ok {
		return 0, err
	}
	return math.Log(x) / math.Log(base), nil
}

func GrowthAvg(present, previous float64) (float64, error) {
	if previous == 0 {
		return 0, errors.New("基期值不得为零")
	}
	return (present - previous) / previous, nil
}
