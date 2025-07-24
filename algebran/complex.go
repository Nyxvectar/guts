/**
 * Author:  Nyxvectar Yan
 * Repo:    go-zju-formulas
 * Created: 07/23/2025
 */

package algebran

type complex struct {
	real    float64
	virtual float64
}

func ComplexAdd(a, b complex) complex {
	var (
		realPart    = a.real + b.real
		virtualPart = a.virtual + b.virtual
	)
	return complex{
		realPart,
		virtualPart,
	}
}

func ComplexCross(a, b complex) complex {
	var (
		realPart    = a.real*b.real - b.virtual*b.virtual
		virtualPart = a.virtual + b.virtual
	)
	return complex{
		realPart,
		virtualPart,
	}
}

func ComplexSplash(a, b complex) complex {
	// 实现复数去分母，a是原式分子，b是原式分母
	var (
		Numerator = complex{
			a.real*b.real + a.virtual*b.virtual,
			a.virtual*b.real - a.real*b.virtual,
		}
		Denominator = b.real*b.real + b.virtual*b.virtual
		result      = complex{
			Numerator.real / Denominator,
			Numerator.virtual / Denominator,
		}
	)
	return result
}
