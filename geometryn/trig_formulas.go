/**
 * Author:  Nyxvectar Yan
 * Repo:    go-zju-formulas
 * Created: 07/23/2025
 */

//"和差化积公式",
//"积化和差公式",
//"三角万能公式",
//"辅助角公式",
//"三角函数周期公式",
//"常见三角不等式"

package geometryn

import (
	"math"
)

func angleToRadian(angle float64) float64 {
	radian := angle * math.Pi / 180
	return radian
}

func ConvertCS(value float64) float64 {
	return math.Sqrt(1 - value*value)
}

func SinSum(sinA, sinB float64) float64 {
	return sinA*ConvertCS(sinB) + sinB*ConvertCS(sinA)
}

func SinDiff(sinA, sinB float64) float64 {
	return sinA*ConvertCS(sinB) - sinB*ConvertCS(sinA)
}

func CosSum(cosA, cosB float64) float64 {
	return cosA*cosB - ConvertCS(cosA)*ConvertCS(cosB)
}

func CosDiff(cosA, cosB float64) float64 {
	return cosA*cosB + ConvertCS(cosA)*ConvertCS(cosB)
}

func SinMulti(sin float64) float64 {
	return 2 * sin * ConvertCS(sin)
}

func CosMulti(cos float64) float64 {
	return 2*cos*cos - 1
}

func TanMulti(tan float64) float64 {
	return (2 * tan) / (1 - tan*tan)
}
