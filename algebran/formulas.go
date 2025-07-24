/**
 * Author:  Nyxvectar Yan
 * Repo:    go-zju-formulas
 * Created: 07/23/2025
 */
// Package algebran 提供代数运算的基本功能
// 导入路径: "go-zju-formulas/algebran"
package algebran

import (
	"errors"
	"math"
)

// CubicDiff 计算两数的立方差
// 返回值为 a³ - b³
// 计算公式：a³ - b³ = (a - b)(a² + ab + b²)
func CubicDiff(a, b float64) float64 {
	return (a - b) * (a*a + a*b + b*b)
}

// SubsetsCount 计算集合的子集总数
// 参数 omnibus 表示集合的元素个数
// 返回值为 2^n，其中 n 是集合的元素数量
// 注意：当 n 大于 64 时可能会导致溢出
func SubsetsCount(omnibus uint64) uint64 {
	return uint64(math.Pow(2, float64(omnibus)))
}

// MeanInequality 计算一组正数的四种平均数并验证均值不等式
// 返回值依次为：调和平均数、几何平均数、算术平均数、平方平均数
// 满足数学关系：H ≤ G ≤ A ≤ Q
// 参数要求：
//   - 输入不能为空集
//   - 所有元素必须为正数
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
		productTotal   float64 = 1 // 初始化乘积为1
		squareTotal    float64
		plusTotal      float64
	)

	for _, i := range u {
		fractionsTotal += 1 / i
		productTotal *= i
		squareTotal += i * i
		plusTotal += i
	}

	// 四种平均数的计算公式
	H := length / fractionsTotal          // 调和平均数
	G := math.Pow(productTotal, 1/length) // 几何平均数
	A := plusTotal / length               // 算术平均数
	Q := math.Sqrt(squareTotal / length)  // 平方平均数

	return H, G, A, Q, nil
}

// CauchyInequality 计算柯西不等式取等条件下的最小值
// 柯西不等式: (a² + b²)(c² + d²) ≥ (ac + bd)²
// 当且仅当 ad = bc 时取等号
// 返回值为 (ac + bd)² 的最小值
// 参数要求：必须满足 a*d == b*c，否则返回错误
func CauchyInequality(a, b, c, d float64) (float64, error) {
	if a*d != b*c {
		return 0, errors.New("实参不满足柯西求最值条件")
	}

	return math.Pow(a*c+b*d, 2), nil
}
