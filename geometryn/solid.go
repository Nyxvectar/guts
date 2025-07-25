/**
 * Author:  Nyxvectar Yan
 * Repo:    go-zju-formulas
 * Created: 07/23/2025
 */

//"祖暅原理",
//"球体表面积与体积公式",
//"欧拉定理"

package geometryn

import (
	"errors"
	"math"
)

var (
	dntExist = "不存在这样的半径/高/母线/底面积"
)

func checker(a, b float64) bool {
	if a <= 0 || b <= 0 {
		return false
	} else {
		return true
	}
}

func AreaCylinder(r, h float64) (float64, error) {
	if checker(r, h) {
		return 2 * math.Pi * r * (r + h), nil
	} else {
		return 0, errors.New(dntExist)
	}
}

func AreaCone(r, l float64) (float64, error) {
	if checker(r, l) {
		return math.Pi * r * (r + l), nil
	} else {
		return 0, errors.New(dntExist)
	}
}

func AreaFlatTopCone(r1, r2, l float64) (float64, error) {
	if checker(r1, l) && r2 > 0 {
		return math.Pi * (r1*r1 + r2*r2 + (r1+r2)*l), nil
	} else {
		return 0, errors.New(dntExist)
	}
}

func VolumeCylinder(s, h float64) (float64, error) {
	if checker(s, h) {
		return s * h, nil
	} else {
		return 0, errors.New(dntExist)
	}
}

func VolumeCone(s, h float64) (float64, error) {
	if checker(s, h) {
		return s * h / 3, nil
	} else {
		return 0, errors.New(dntExist)
	}
}

func VolumeFlatTopCone(s1, s2, h float64) (float64, error) {
	if checker(s1, h) && s2 > 0 {
		return (s1 + s2 + math.Sqrt(s1*s2)) * h / 3, nil
	} else {
		return 0, errors.New(dntExist)
	}
}
