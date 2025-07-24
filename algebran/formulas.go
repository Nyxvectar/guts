/**
 * Author:  Nyxvectar Yan
 * Repo:    go-zju-formulas
 * Created: 07/23/2025
 */

//"浓度不等式",
//"绝对值三角不等式",
//"柯西不等式"

package algebran

import (
	"errors"
	"math"
)

func CubicDiff(a, b float64) float64 {
	return (a - b) * (a*a + a*b + b*b)
}

func SubsetsCount(omnibus uint64) uint64 {
	var subsets = uint64(math.Pow(2, float64(omnibus)))
	return subsets
}

func MeanInequality(u []float64) (float64, float64, float64, float64, error) {
	for _, num := range u {
		if num > 0 != true {
			var errNew = "集合内需全部为正数"
			return 0, 0, 0, 0, errors.New(errNew)
		}
	}
	var length = float64(len(u))
	{
		var (
			H              float64
			G              float64
			A              float64
			Q              float64
			fractionsTotal float64
			productTotal   float64
			squareTotal    float64
			plusTotal      float64
		)

		for _, i := range u {
			fractionsTotal += 1 / i
		}
		for _, i := range u {
			productTotal *= i
		}
		for _, i := range u {
			squareTotal += i * i
		}
		for _, i := range u {
			plusTotal += i
		}

		H = length / fractionsTotal           //调和平均数
		G = math.Pow(productTotal, 1/length)  //几何平均数
		A = plusTotal / length                //算术平均数
		Q = math.Pow(squareTotal/length, 1/2) //平方平均数

		// 始终有 H<=G<=A<=Q 成立，按递增次序return。
		return H, G, A, Q, nil
	}
}
