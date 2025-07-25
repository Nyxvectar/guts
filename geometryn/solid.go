/**
 * Author:  Nyxvectar Yan
 * Repo:    go-zju-formulas
 * Created: 07/23/2025
 */

package geometryn

import (
	"errors"
	"math"
)

var (
	dntExist = "不存在这样的几何体"
)

func checker(a, b float64) bool {
	if a < 0 || b < 0 {
		return false
	} else {
		return true
	}
}

// 注意到台体的公式事实上是通用于圆柱体和圆锥的
// 即后二者的公式只是台体公式的特殊情况，那么我
// 们就可以只使用台体公式来同时解决此三者的问题
// 对于圆锥，认为她的其中一个地面半径和面积都为
// 零即可；对于圆柱体，认为其上下面是相同的即可

func AreaCycle(r, h float64) (float64, error) {
	if checker(r, h) {
		return 2 * math.Pi * r * (r + h), nil
	} else {
		return 0, errors.New(dntExist)
	}
}

func VolumeCycle(s1, s2, h float64) (float64, error) {
	if checker(s1, h) && s2 > 0 {
		return (s1 + s2 + math.Sqrt(s1*s2)) * h / 3, nil
	} else {
		return 0, errors.New(dntExist)
	}
}

func AreaSphere(r float64) (float64, error) {
	if checker(r, 1) {
		return 4 * math.Pi * r * r, nil
	} else {
		return 0, errors.New(dntExist)
	}
}

func VolumeSphere(r float64) (float64, error) {
	if checker(r, 1) {
		return 4 * math.Pi * r * r * r / 3, nil
	} else {
		return 0, errors.New(dntExist)
	}
}

func EulerTheorem(vertices, edges, faces uint64) (uint64, error) {
	// 利用欧拉定理，待求的值设置为零即可返回该值的
	// 欲求结果，或者全部传入以判断该几何体是否成立
	if vertices == 0 {
		if 2-faces+edges >= 0 {
			return 2 - faces + edges, nil
		} else {
			return 0, errors.New(dntExist)
		}
	} else if edges == 0 {
		if vertices+faces-2 >= 0 {
			return vertices + faces - 2, nil
		} else {
			return 0, errors.New(dntExist)
		}
	} else if faces == 0 {
		if 2+edges-vertices >= 0 {
			return 2 + edges - vertices, nil
		} else {
			return 0, errors.New(dntExist)
		}
	} else {
		if vertices-edges+faces == 2 {
			return 0, nil
		} else {
			return 0, errors.New(dntExist)
		}
	}
}
