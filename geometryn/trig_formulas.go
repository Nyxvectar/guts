/**
 * Author:  Nyxvectar Yan
 * Repo:    go-zju-formulas
 * Created: 07/23/2025
 */

// Package geometryn 提供三角函数及相关数学转换、公式的实现，包括基本三角函数、角度弧度转换、和角公式、辅助角公式等
package geometryn

import (
	"errors"
	"math"
)

var (
	outDefinition = "函数在此时没有定义"       // 函数无定义时的错误信息（如正切在π/2+kπ处）
	outRange      = "正/余弦值超出范围[-1, 1]" // 输入值超出[-1,1]范围时的错误信息
	outRule       = "Omega不得为零"            // 角频率为零时的错误信息（周期计算中）
)

// isInRange 检查值是否在[min, max]闭区间内
// 参数：
//
//	value：待检查的值
//	min：区间最小值
//	max：区间最大值
//
// 返回值：
//
//	若value在[min, max]范围内则返回true，否则返回false
func isInRange(value, min, max float64) bool {
	return value >= min && value <= max
}

// Sin 计算角度（弧度）的正弦值
// 参数：
//
//	rad：角度（弧度制）
//
// 返回值：
//
//	该角度的正弦值（范围[-1, 1]）
func Sin(rad float64) float64 {
	return math.Sin(rad)
}

// Cos 计算角度（弧度）的余弦值
// 参数：
//
//	rad：角度（弧度制）
//
// 返回值：
//
//	该角度的余弦值（范围[-1, 1]）
func Cos(rad float64) float64 {
	return math.Cos(rad)
}

// Tan 计算角度（弧度）的正切值
// 参数：
//
//	rad：角度（弧度制）
//
// 返回值：
//
//	该角度的正切值；若角度的余弦值接近0（即角度为π/2 + kπ），返回错误
//
// 错误情况：
//   - 当cos(rad)≈0时，正切无定义（outDefinition）
func Tan(rad float64) (float64, error) {
	if math.Abs(math.Cos(rad)) < 1e-10 {
		return 0, errors.New(outDefinition)
	}
	return math.Tan(rad), nil
}

// DegToRad 将角度（度）转换为弧度
// 参数：
//
//	deg：角度（度制，如90表示90度）
//
// 返回值：
//
//	对应的弧度值（公式：弧度 = 度 × π / 180）
func DegToRad(deg float64) float64 {
	return deg * math.Pi / 180
}

// RadToDeg 将弧度转换为角度（度）
// 参数：
//
//	rad：角度（弧度制，如π/2）
//
// 返回值：
//
//	对应的度值（公式：度 = 弧度 × 180 / π）
func RadToDeg(rad float64) float64 {
	return rad * 180 / math.Pi
}

// SinToCos 由正弦值计算对应的余弦值（取非负根）
// 参数：
//
//	sin：正弦值（需在[-1, 1]范围内）
//
// 返回值：
//
//	对应的余弦值（√(1 - sin²)）；若sin超出范围，返回错误
//
// 错误情况：
//   - sin∉[-1, 1]时（outRange）
func SinToCos(sin float64) (float64, error) {
	if !isInRange(sin, -1, 1) {
		return 0, errors.New(outRange)
	}
	return math.Sqrt(1 - sin*sin), nil
}

// CosToSin 由余弦值计算对应的正弦值（取非负根）
// 参数：
//
//	cos：余弦值（需在[-1, 1]范围内）
//
// 返回值：
//
//	对应的正弦值（√(1 - cos²)）；若cos超出范围，返回错误
//
// 错误情况：
//   - cos∉[-1, 1]时（outRange）
func CosToSin(cos float64) (float64, error) {
	if !isInRange(cos, -1, 1) {
		return 0, errors.New(outRange)
	}
	return math.Sqrt(1 - cos*cos), nil
}

// SinSum 计算两角和的正弦值（和角公式）
// 公式：sin(A + B) = sinA·cosB + cosA·sinB
// 参数：
//
//	radA：角A（弧度）
//	radB：角B（弧度）
//
// 返回值：
//
//	sin(A + B)的值
func SinSum(radA, radB float64) float64 {
	return Sin(radA)*Cos(radB) + Cos(radA)*Sin(radB)
}

// SinDiff 计算两角差的正弦值（差角公式）
// 公式：sin(A - B) = sinA·cosB - cosA·sinB
// 参数：
//
//	radA：角A（弧度）
//	radB：角B（弧度）
//
// 返回值：
//
//	sin(A - B)的值
func SinDiff(radA, radB float64) float64 {
	return Sin(radA)*Cos(radB) - Cos(radA)*Sin(radB)
}

// CosSum 计算两角和的余弦值（和角公式）
// 公式：cos(A + B) = cosA·cosB - sinA·sinB
// 参数：
//
//	radA：角A（弧度）
//	radB：角B（弧度）
//
// 返回值：
//
//	cos(A + B)的值
func CosSum(radA, radB float64) float64 {
	return Cos(radA)*Cos(radB) - Sin(radA)*Sin(radB)
}

// CosDiff 计算两角差的余弦值（差角公式）
// 公式：cos(A - B) = cosA·cosB + sinA·sinB
// 参数：
//
//	radA：角A（弧度）
//	radB：角B（弧度）
//
// 返回值：
//
//	cos(A - B)的值
func CosDiff(radA, radB float64) float64 {
	return Cos(radA)*Cos(radB) + Sin(radA)*Sin(radB)
}

// SinDoubleAngle 计算二倍角的正弦值（二倍角公式）
// 公式：sin(2θ) = 2sinθ·cosθ
// 参数：
//
//	rad：角θ（弧度）
//
// 返回值：
//
//	sin(2θ)的值
func SinDoubleAngle(rad float64) float64 {
	return 2 * Sin(rad) * Cos(rad)
}

// CosDoubleAngle 计算二倍角的余弦值（二倍角公式）
// 公式：cos(2θ) = 2cos²θ - 1
// 参数：
//
//	rad：角θ（弧度）
//
// 返回值：
//
//	cos(2θ)的值
func CosDoubleAngle(rad float64) float64 {
	return 2*Cos(rad)*Cos(rad) - 1
}

// TanDoubleAngle 计算二倍角的正切值（二倍角公式）
// 公式：tan(2θ) = 2tanθ / (1 - tan²θ)
// 参数：
//
//	rad：角θ（弧度）
//
// 返回值：
//
//	tan(2θ)的值；若tanθ无定义或分母为0，返回错误
//
// 错误情况：
//   - 当θ的正切无定义时（outDefinition，同Tan函数）
//   - 当1 - tan²θ≈0时（即θ≈π/4 + kπ/2），无定义（outDefinition）
func TanDoubleAngle(rad float64) (float64, error) {
	tan, err := Tan(rad)
	if err != nil {
		return 0, err
	}
	if math.Abs(1-tan*tan) < 1e-10 {
		return 0, errors.New(outDefinition)
	}
	return (2 * tan) / (1 - tan*tan), nil
}

// SumToProductSinSum 正弦和的和差化积（和差化积公式）
// 公式：sinA + sinB = 2sin[(A+B)/2]·cos[(A-B)/2]
// 参数：
//
//	radA：角A（弧度）
//	radB：角B（弧度）
//
// 返回值：
//
//	sinA + sinB的和差化积结果
func SumToProductSinSum(radA, radB float64) (float64, float64) {
	return 2 * Sin((radA+radB)/2) * Cos((radA-radB)/2), 0
}

// SumToProductSinDiff 正弦差的和差化积（和差化积公式）
// 公式：sinA - sinB = 2cos[(A+B)/2]·sin[(A-B)/2]
// 参数：
//
//	radA：角A（弧度）
//	radB：角B（弧度）
//
// 返回值：
//
//	sinA - sinB的和差化积结果
func SumToProductSinDiff(radA, radB float64) (float64, float64) {
	return 2 * Cos((radA+radB)/2) * Sin((radA-radB)/2), 0
}

// SumToProductCosSum 余弦和的和差化积（和差化积公式）
// 公式：cosA + cosB = 2cos[(A+B)/2]·cos[(A-B)/2]
// 参数：
//
//	radA：角A（弧度）
//	radB：角B（弧度）
//
// 返回值：
//
//	cosA + cosB的和差化积结果
func SumToProductCosSum(radA, radB float64) (float64, float64) {
	return 2 * Cos((radA+radB)/2) * Cos((radA-radB)/2), 0
}

// SumToProductCosDiff 余弦差的和差化积（和差化积公式）
// 公式：cosA - cosB = -2sin[(A+B)/2]·sin[(A-B)/2]
// 参数：
//
//	radA：角A（弧度）
//	radB：角B（弧度）
//
// 返回值：
//
//	cosA - cosB的和差化积结果
func SumToProductCosDiff(radA, radB float64) (float64, float64) {
	return -2 * Sin((radA+radB)/2) * Sin((radA-radB)/2), 0
}

// ProductToSumSinCos 正弦余弦积的积化和差（积化和差公式）
// 公式：sinA·cosB = [sin(A+B) + sin(A-B)] / 2
// 参数：
//
//	radA：角A（弧度）
//	radB：角B（弧度）
//
// 返回值：
//
//	两个结果分别为[sin(A+B)]/2和[sin(A-B)]/2
func ProductToSumSinCos(radA, radB float64) (float64, float64) {
	return 0.5 * Sin(radA+radB), 0.5 * Sin(radA-radB)
}

// ProductToSumSinSin 正弦正弦积的积化和差（积化和差公式）
// 公式：sinA·sinB = [cos(A-B) - cos(A+B)] / 2
// 参数：
//
//	radA：角A（弧度）
//	radB：角B（弧度）
//
// 返回值：
//
//	两个结果分别为[cos(A-B)]/2和[-cos(A+B)]/2
func ProductToSumSinSin(radA, radB float64) (float64, float64) {
	return 0.5 * Cos(radA-radB), -0.5 * Cos(radA+radB)
}

// ProductToSumCosCos 余弦余弦积的积化和差（积化和差公式）
// 公式：cosA·cosB = [cos(A-B) + cos(A+B)] / 2
// 参数：
//
//	radA：角A（弧度）
//	radB：角B（弧度）
//
// 返回值：
//
//	两个结果分别为[cos(A-B)]/2和[cos(A+B)]/2
func ProductToSumCosCos(radA, radB float64) (float64, float64) {
	return 0.5 * Cos(radA-radB), 0.5 * Cos(radA+radB)
}

// SinHalfAngle 半角的正弦值（半角公式，取非负根）
// 公式：sin(θ/2) = √[(1 - cosθ)/2]
// 参数：
//
//	cos：cosθ的值（需在[-1, 1]范围内）
//
// 返回值：
//
//	sin(θ/2)的值；若cos超出范围，返回错误
//
// 错误情况：
//   - cos∉[-1, 1]时（outRange）
func SinHalfAngle(cos float64) (float64, error) {
	if !isInRange(cos, -1, 1) {
		return 0, errors.New(outRange)
	}
	return math.Sqrt((1 - cos) / 2), nil
}

// CosHalfAngle 半角的余弦值（半角公式，取非负根）
// 公式：cos(θ/2) = √[(1 + cosθ)/2]
// 参数：
//
//	cos：cosθ的值（需在[-1, 1]范围内）
//
// 返回值：
//
//	cos(θ/2)的值；若cos超出范围，返回错误
//
// 错误情况：
//   - cos∉[-1, 1]时（outRange）
func CosHalfAngle(cos float64) (float64, error) {
	if !isInRange(cos, -1, 1) {
		return 0, errors.New(outRange)
	}
	return math.Sqrt((1 + cos) / 2), nil
}

// TanHalfAngle 半角的正切值（半角公式，取非负根）
// 公式：tan(θ/2) = √[(1 - cosθ)/(1 + cosθ)]
// 参数：
//
//	cos：cosθ的值（需在[-1, 1]范围内）
//
// 返回值：
//
//	tan(θ/2)的值；若cos超出范围或分母为0，返回错误
//
// 错误情况：
//   - cos∉[-1, 1]时（outRange）
//   - 当1 + cosθ≈0时（即θ≈π + 2kπ），无定义（outDefinition）
func TanHalfAngle(cos float64) (float64, error) {
	if !isInRange(cos, -1, 1) {
		return 0, errors.New(outRange)
	}
	if math.Abs(1+cos) < 1e-10 {
		return 0, errors.New(outDefinition)
	}
	return math.Sqrt((1 - cos) / (1 + cos)), nil
}

// SinFromHalf 由半角正切表示全角正弦（万能公式）
// 公式：sinθ = 2tan(θ/2) / (1 + tan²(θ/2))
// 参数：
//
//	tan：tan(θ/2)的值
//
// 返回值：
//
//	sinθ的值
func SinFromHalf(tan float64) float64 {
	return 2 * tan / (1 + tan*tan)
}

// CosFromHalf 由半角正切表示全角余弦（万能公式）
// 公式：cosθ = (1 - tan²(θ/2)) / (1 + tan²(θ/2))
// 参数：
//
//	tan：tan(θ/2)的值
//
// 返回值：
//
//	cosθ的值
func CosFromHalf(tan float64) float64 {
	return (1 - tan*tan) / (1 + tan*tan)
}

// TanFromHalf 由半角正切表示全角正切（万能公式）
// 公式：tanθ = 2tan(θ/2) / (1 - tan²(θ/2))
// 参数：
//
//	tan：tan(θ/2)的值
//
// 返回值：
//
//	tanθ的值
func TanFromHalf(tan float64) float64 {
	return 2 * tan / (1 - tan*tan)
}

// AuxiliaryAngle 辅助角公式（将a·sinx + b·cosx化为A·sin(x + y)）
// 公式：a·sinx + b·cosx = A·sin(x + y)，其中A=√(a²+b²)，y=arctan(b/a)
// 参数：
//
//	a, b：系数（a和b不同时为0）
//
// 返回值：
//
//	A（振幅，√(a²+b²)）和y（相位，arctan(b/a)）；若a和b都为0，返回错误
//
// 错误情况：
//   - 当a=0且b=0时，无定义（outDefinition）
func AuxiliaryAngle(a, b float64) (float64, float64, error) {
	if a == 0 && b == 0 {
		return 0, 0, errors.New(outDefinition)
	}
	A := math.Sqrt(a*a + b*b)
	y := math.Atan2(b, a)
	return A, y, nil
}

// InverseAuxiliaryAngle 辅助角公式的逆过程（由A和y得到a和b）
// 公式：A·sin(x + y) = a·sinx + b·cosx，其中a=A·cosy，b=A·siny
// 参数：
//
//	A：振幅
//	y：相位（弧度）
//
// 返回值：
//
//	系数a和b（a=A·cosy，b=A·siny）
func InverseAuxiliaryAngle(A, y float64) (float64, float64) {
	a := A * math.Cos(y)
	b := A * math.Sin(y)
	return a, b
}

// GetTerm 计算周期（由角频率w）
// 公式：周期T = 2π / w（w为角频率，w≠0）
// 参数：
//
//	w：角频率（w≠0）
//
// 返回值：
//
//	周期T；若w=0，返回错误
//
// 错误情况：
//   - 当w=0时，无定义（outRule）
func GetTerm(w float64) (float64, error) {
	if w == 0 {
		return 0, errors.New(outRule)
	}
	return 2 * math.Pi / w, nil
}
