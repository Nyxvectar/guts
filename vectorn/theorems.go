/**
 * Author:  Nyxvectar Yan
 * Repo:    go-zju-formulas
 * Created: 07/23/2025
 */

// Package vectorn 提供向量运算的基本功能
// 导入路径: "go-zju-formulas/algebran"
package vectorn

import "math"

// Vector 表示三维空间中的向量
type Vector struct {
	X float64 // X轴分量
	Y float64 // Y轴分量
	Z float64 // Z轴分量
}

// VectorModulus 计算向量的模长（长度）
// 返回值为向量的模长，即 √(X² + Y² + Z²)
// 使用math.Hypot两次嵌套计算，避免数值溢出
func VectorModulus(v Vector) float64 {
	return math.Hypot(v.X, math.Hypot(v.Y, v.Z))
}

// DeterminantLinear 判断两个向量是否线性相关
// 零向量与任何向量都线性相关
// 通过检查坐标分量的交叉乘积是否近似相等来判断
func DeterminantLinear(a, b Vector) bool {
	if isZeroVector(a) || isZeroVector(b) {
		return true
	}
	if !almostEqual(a.X*b.Y, a.Y*b.X) ||
		!almostEqual(a.X*b.Z, a.Z*b.X) ||
		!almostEqual(a.Y*b.Z, a.Z*b.Y) {
		return false
	}
	return true
}

// QuantityProduct 计算两个向量的数量积（点积）
// 返回值为 a·b = a.X*b.X + a.Y*b.Y + a.Z*b.Z
func QuantityProduct(a, b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// VectorAngleCos 计算两个向量夹角的余弦值
// 若任一向量为零向量，返回0
// 计算公式为 cosθ = (a·b) / (|a|*|b|)
func VectorAngleCos(a, b Vector) float64 {
	if isZeroVector(a) || isZeroVector(b) {
		return 0
	}
	return QuantityProduct(a, b) / (VectorModulus(a) * VectorModulus(b))
}

// CrossProduct 计算两个向量的叉积（向量积）
// 返回向量垂直于a和b构成的平面，方向遵循右手定则
// 计算公式为：
//
//	X = a.Y*b.Z - a.Z*b.Y
//	Y = a.Z*b.X - a.X*b.Z
//	Z = a.X*b.Y - a.Y*b.X
func CrossProduct(a, b Vector) Vector {
	return Vector{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
	}
}

// isZeroVector 判断向量是否近似为零向量
// 使用1e-10的容差进行浮点数比较
func isZeroVector(v Vector) bool {
	const epsilon = 1e-10
	return math.Abs(v.X) < epsilon &&
		math.Abs(v.Y) < epsilon &&
		math.Abs(v.Z) < epsilon
}

// almostEqual 判断两个浮点数是否近似相等
// 使用绝对误差和相对误差结合的比较方法
// 适用于不同数量级的浮点数比较
func almostEqual(a, b float64) bool {
	const epsilon = 1e-10
	return math.Abs(a-b) < epsilon ||
		math.Abs(a-b) < epsilon*math.Max(math.Abs(a), math.Abs(b))
}
