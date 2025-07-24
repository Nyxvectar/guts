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

func CubicDiff(a, b float64) float64 {
	return (a - b) * (a*a + a*b + b*b)
}

func SubsetsCount(omnibus uint64) uint64 {
	return uint64(math.Pow(2, float64(omnibus)))
}

func MeanInequality(u []float64) (float64, float64, float64, float64, error) {
	if len(u) == 0 {
		return 0, 0, 0, 0, errors.New("不得传入空集")
	}
	for _, num := range u {
		if num <= 0 {
			return 0, 0, 0, 0, errors.New("集合内需全部为正数")
		}
	}
	length := float64(len(u))
	var (
		fractionsTotal float64
		productTotal   float64 = 1
		squareTotal    float64
		plusTotal      float64
	)
	for _, i := range u {
		fractionsTotal += 1 / i
		productTotal *= i
		squareTotal += i * i
		plusTotal += i
	}
	H := length / fractionsTotal          // 调和平均数
	G := math.Pow(productTotal, 1/length) // 几何平均数
	A := plusTotal / length               // 算术平均数
	Q := math.Sqrt(squareTotal / length)  // 平方平均数
	return H, G, A, Q, nil
}

func CauchyInequality(a, b, c, d float64) (float64, error) {
	if a*d != b*c {
		return 0, errors.New("实参不满足柯西求最值条件")
	}
	return math.Pow(a*c+b*d, 2), nil
}
