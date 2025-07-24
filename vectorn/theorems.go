/**
 * Author:  Nyxvectar Yan
 * Repo:    go-zju-formulas
 * Created: 07/23/2025
 */

package vectorn

import "math"

type Vector struct {
	X float64
	Y float64
	Z float64
}

func VectorModulus(v Vector) float64 {
	return math.Hypot(v.X, math.Hypot(v.Y, v.Z))
}

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

func QuantityProduct(a, b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func VectorAngleCos(a, b Vector) float64 {
	if isZeroVector(a) || isZeroVector(b) {
		return 0
	}
	return QuantityProduct(a, b) / (VectorModulus(a) * VectorModulus(b))
}

func CrossProduct(a, b Vector) Vector {
	return Vector{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
	}
}

func isZeroVector(v Vector) bool {
	const epsilon = 1e-10
	return math.Abs(v.X) < epsilon &&
		math.Abs(v.Y) < epsilon &&
		math.Abs(v.Z) < epsilon
}

func almostEqual(a, b float64) bool {
	const epsilon = 1e-10
	return math.Abs(a-b) < epsilon ||
		math.Abs(a-b) < epsilon*math.Max(math.Abs(a), math.Abs(b))
}
