/**
 * Author:  Nyxvectar Yan
 * Repo:    go-zju-formulas
 * Created: 07/23/2025
 */

package vectorn

import "math"

type Vector struct {
	x float64
	y float64
	z float64
}

func VectorModulus(vectar Vector) float64 {
	return math.Sqrt(vectar.x*vectar.x + vectar.y*vectar.y + vectar.z*vectar.z)
}

func DeterminantLinear(a, b Vector) bool {
	if (a.x == 0 && a.y == 0 && a.z == 0) || (b.x == 0 && b.y == 0 && b.z == 0) {
		return true
	} else if a.x*b.y == b.x*a.y && a.x*b.z == b.x*a.z && a.y*b.z == b.y*a.z {
		return true
	} else {
		return false
	}
}

func QuantityProduct(a, b Vector) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func VectorAngleCos(a, b Vector) float64 {
	return QuantityProduct(a, b) / (VectorModulus(a) * VectorModulus(b))
}

func CrossProduct(a, b Vector) Vector {
	return Vector{
		a.y*b.z - a.z*b.y,
		a.z*b.x - a.x*b.z,
		a.x*b.y - a.y*b.x,
	}
}
