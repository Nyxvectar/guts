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
	outDefinition = "函数在此时没有定义"
	outRange      = "正/余弦值超出范围[-1, 1]"
	outRule       = "Omega不得为零"
)

func isInRange(value, min, max float64) bool {
	return value >= min && value <= max
}

func Sin(rad float64) float64 {
	return math.Sin(rad)
}

func Cos(rad float64) float64 {
	return math.Cos(rad)
}

func Tan(rad float64) (float64, error) {
	if math.Abs(math.Cos(rad)) < 1e-10 {
		return 0, errors.New(outDefinition)
	}
	return math.Tan(rad), nil
}

func DegToRad(deg float64) float64 {
	return deg * math.Pi / 180
}

func RadToDeg(rad float64) float64 {
	return rad * 180 / math.Pi
}

func SinToCos(sin float64) (float64, error) {
	if !isInRange(sin, -1, 1) {
		return 0, errors.New(outRange)
	}
	return math.Sqrt(1 - sin*sin), nil
}

func CosToSin(cos float64) (float64, error) {
	if !isInRange(cos, -1, 1) {
		return 0, errors.New(outRange)
	}
	return math.Sqrt(1 - cos*cos), nil
}

func SinSum(radA, radB float64) float64 {
	return Sin(radA)*Cos(radB) + Cos(radA)*Sin(radB)
}

func SinDiff(radA, radB float64) float64 {
	return Sin(radA)*Cos(radB) - Cos(radA)*Sin(radB)
}

func CosSum(radA, radB float64) float64 {
	return Cos(radA)*Cos(radB) - Sin(radA)*Sin(radB)
}

func CosDiff(radA, radB float64) float64 {
	return Cos(radA)*Cos(radB) + Sin(radA)*Sin(radB)
}

func SinDoubleAngle(rad float64) float64 {
	return 2 * Sin(rad) * Cos(rad)
}

func CosDoubleAngle(rad float64) float64 {
	return 2*Cos(rad)*Cos(rad) - 1
}

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

func SumToProductSinSum(radA, radB float64) (float64, float64) {
	return 2 * Sin((radA+radB)/2) * Cos((radA-radB)/2), 0
}

func SumToProductSinDiff(radA, radB float64) (float64, float64) {
	return 2 * Cos((radA+radB)/2) * Sin((radA-radB)/2), 0
}

func SumToProductCosSum(radA, radB float64) (float64, float64) {
	return 2 * Cos((radA+radB)/2) * Cos((radA-radB)/2), 0
}

func SumToProductCosDiff(radA, radB float64) (float64, float64) {
	return -2 * Sin((radA+radB)/2) * Sin((radA-radB)/2), 0
}

func ProductToSumSinCos(radA, radB float64) (float64, float64) {
	return 0.5 * Sin(radA+radB), 0.5 * Sin(radA-radB)
}

func ProductToSumSinSin(radA, radB float64) (float64, float64) {
	return 0.5 * Cos(radA-radB), -0.5 * Cos(radA+radB)
}

func ProductToSumCosCos(radA, radB float64) (float64, float64) {
	return 0.5 * Cos(radA-radB), 0.5 * Cos(radA+radB)
}

func SinHalfAngle(cos float64) (float64, error) {
	if !isInRange(cos, -1, 1) {
		return 0, errors.New(outRange)
	}
	return math.Sqrt((1 - cos) / 2), nil
}

func CosHalfAngle(cos float64) (float64, error) {
	if !isInRange(cos, -1, 1) {
		return 0, errors.New(outRange)
	}
	return math.Sqrt((1 + cos) / 2), nil
}

func TanHalfAngle(cos float64) (float64, error) {
	if !isInRange(cos, -1, 1) {
		return 0, errors.New(outRange)
	}
	if math.Abs(1+cos) < 1e-10 {
		return 0, errors.New(outDefinition)
	}
	return math.Sqrt((1 - cos) / (1 + cos)), nil
}

func SinFromHalf(tan float64) float64 {
	return 2 * tan / (1 + tan*tan)
}

func CosFromHalf(tan float64) float64 {
	return (1 - tan*tan) / (1 + tan*tan)
}

func TanFromHalf(tan float64) float64 {
	return 2 * tan / (1 - tan*tan)
}

func AuxiliaryAngle(a, b float64) (float64, float64, error) {
	if a == 0 && b == 0 {
		return 0, 0, errors.New(outDefinition)
	}
	A := math.Sqrt(a*a + b*b)
	y := math.Atan2(b, a)
	return A, y, nil
}

func InverseAuxiliaryAngle(A, y float64) (float64, float64) {
	a := A * math.Cos(y)
	b := A * math.Sin(y)
	return a, b
}

func GetTerm(w float64) (float64, error) {
	if w == 0 {
		return 0, errors.New(outRule)
	}
	return 2 * math.Pi / w, nil
}
