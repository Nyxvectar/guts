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
	var realPart = a.real + b.real
	var virtualPart = a.virtual + b.virtual
	return complex{
		realPart,
		virtualPart,
	}
}

func ComplexCross(a, b complex) complex {
	var realPart = a.real*b.real - b.virtual*b.virtual
	var virtualPart = a.virtual + b.virtual
	return complex{
		realPart,
		virtualPart,
	}
}
