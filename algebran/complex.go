/**
 * Author:  Nyxvectar Yan
 * Repo:    go-zju-formulas
 * Created: 07/23/2025
 */

// Package algebran 提供复数运算的基本功能
// 导入路径: "go-zju-formulas/algebran"
package algebran

import "math"

// Complex 表示复数，由实部和虚部组成
// 复数的标准形式为 a+bi，其中a为实部，b为虚部
type Complex struct {
	Real      float64 // 实部
	Imaginary float64 // 虚部（使用Imaginary替代virtual更符合数学术语）
}

// ComplexAdd 计算两个复数的和
// 返回值为 a+b，其中 a=a₁+a₂i，b=b₁+b₂i
// 计算公式：(a₁+b₁) + (a₂+b₂)i
func ComplexAdd(a, b Complex) Complex {
	return Complex{
		Real:      a.Real + b.Real,
		Imaginary: a.Imaginary + b.Imaginary,
	}
}

// ComplexMultiply 计算两个复数的乘积
// 返回值为 a×b，其中 a=a₁+a₂i，b=b₁+b₂i
// 计算公式：(a₁×b₁ - a₂×b₂) + (a₁×b₂ + a₂×b₁)i
func ComplexMultiply(a, b Complex) Complex {
	return Complex{
		Real:      a.Real*b.Real - a.Imaginary*b.Imaginary,
		Imaginary: a.Real*b.Imaginary + a.Imaginary*b.Real,
	}
}

// ComplexDivide 计算两个复数的商
// 返回值为 a/b，其中 a=a₁+a₂i，b=b₁+b₂i
// 计算公式：(a×共轭(b))/(|b|²)
// 注意：当分母为0+0i时会触发除零错误
func ComplexDivide(a, b Complex) Complex {
	denominator := b.Real*b.Real + b.Imaginary*b.Imaginary

	// 检查分母是否接近零（考虑浮点数精度）
	if math.Abs(denominator) < 1e-10 {
		panic("complex division by zero")
	}

	return Complex{
		Real:      (a.Real*b.Real + a.Imaginary*b.Imaginary) / denominator,
		Imaginary: (a.Imaginary*b.Real - a.Real*b.Imaginary) / denominator,
	}
}

// ComplexConjugate 计算复数的共轭
// 返回值为 a 的共轭复数，即 a₁-a₂i
func ComplexConjugate(a Complex) Complex {
	return Complex{
		Real:      a.Real,
		Imaginary: -a.Imaginary,
	}
}

// ComplexModulus 计算复数的模（绝对值）
// 返回值为复数的模，即 √(a₁² + a₂²)
func ComplexModulus(a Complex) float64 {
	return math.Sqrt(a.Real*a.Real + a.Imaginary*a.Imaginary)
}
