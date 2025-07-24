/**
 * Author:  Nyxvectar Yan
 * Repo:    go-zju-formulas
 * Created: 07/23/2025
 */

package algebran

import "math"

type Complex struct {
	Real      float64
	Imaginary float64
}

func ComplexAdd(a, b Complex) Complex {
	return Complex{
		Real:      a.Real + b.Real,
		Imaginary: a.Imaginary + b.Imaginary,
	}
}

func ComplexMultiply(a, b Complex) Complex {
	return Complex{
		Real:      a.Real*b.Real - a.Imaginary*b.Imaginary,
		Imaginary: a.Real*b.Imaginary + a.Imaginary*b.Real,
	}
}

func ComplexDivide(a, b Complex) Complex {
	denominator := b.Real*b.Real + b.Imaginary*b.Imaginary
	if math.Abs(denominator) < 1e-10 {
		panic("complex division by zero")
	}
	return Complex{
		Real:      (a.Real*b.Real + a.Imaginary*b.Imaginary) / denominator,
		Imaginary: (a.Imaginary*b.Real - a.Real*b.Imaginary) / denominator,
	}
}

func ComplexConjugate(a Complex) Complex {
	return Complex{
		Real:      a.Real,
		Imaginary: -a.Imaginary,
	}
}

func ComplexModulus(a Complex) float64 {
	return math.Sqrt(a.Real*a.Real + a.Imaginary*a.Imaginary)
}
