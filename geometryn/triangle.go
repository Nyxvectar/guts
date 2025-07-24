/**
 * Author:  Nyxvectar Yan
 * Repo:    go-zju-formulas
 * Created: 07/23/2025
 */

package geometryn

import (
	"errors"
	"math"
)

type Vector2D struct {
	X float64
	Y float64
}

type Triangle struct {
	A Vector2D
	B Vector2D
	C Vector2D
}

var (
	lengthNegative  = "三角形不存在 [length<0]"       
	resultNegative  = "三角形不存在 [square<0]"      
	angleNegative   = "三角形不存在 [angle<0]"      
	angleOutRange   = "三角形不存在 [angleOutRange]"  
	angleFault      = "三角形不存在 [angleTotal!=Pi]" 
	calibrationFail = "计算结果不一 [present!=previous]" 
	calculateFail   = "无法计算     {!}"                
)

func LawOfSines(a, b, c, A, B, C float64) (float64, error) {
	if a <= 0 || b <= 0 || c <= 0 {
		return 0, errors.New(lengthNegative)
	}
	if A <= 0 || B <= 0 || C <= 0 {
		return 0, errors.New(angleNegative)
	}
	if math.Abs(A+B+C-math.Pi) > 1e-9 {
		return 0, errors.New(angleFault)
	}
	R1 := a / math.Sin(A)
	R2 := b / math.Sin(B)
	R3 := c / math.Sin(C)
	if math.Abs(R1-R2) > 1e-9 || math.Abs(R1-R3) > 1e-9 {
		return 0, errors.New(calibrationFail)
	}
	return R1 / 2, nil
}

func LawOfCosines(a, b, C float64) (float64, error) {
	if a <= 0 || b <= 0 {
		return 0, errors.New(lengthNegative)
	}
	if C <= 0 || C >= math.Pi {
		return 0, errors.New(angleOutRange)
	}
	cSquared := a*a + b*b - 2*a*b*math.Cos(C)
	if cSquared < 0 {
		return 0, errors.New(resultNegative)
	}
	return math.Sqrt(cSquared), nil
}

func ProjectionTheorem(a, b, c, B, C float64) (bool, error) {
	if a <= 0 || b <= 0 || c <= 0 {
		return false, errors.New(lengthNegative)
	}
	if B <= 0 || C <= 0 || B+C >= math.Pi {
		return false, errors.New(angleOutRange)
	}
	left := a
	right := b*math.Cos(C) + c*math.Cos(B)
	return math.Abs(left-right) < 1e-9, nil
}

func MedianLength(a, b, c float64) (float64, error) {
	if a <= 0 || b <= 0 || c <= 0 {
		return 0, errors.New(lengthNegative)
	}

	if a >= b+c || b >= a+c || c >= a+b {
		return 0, errors.New(calibrationFail)
	}

	median := math.Sqrt(2*b*b+2*c*c-a*a) / 2
	return median, nil
}

func Centroid(t Triangle) Vector2D {
	return Vector2D{
		X: (t.A.X + t.B.X + t.C.X) / 3,
		Y: (t.A.Y + t.B.Y + t.C.Y) / 3,
	}
}

func Incenter(t Triangle) (Vector2D, error) {
	a := distance(t.B, t.C)
	b := distance(t.A, t.C) 
	c := distance(t.A, t.B) 
	if a <= 0 || b <= 0 || c <= 0 {
		return Vector2D{}, errors.New(lengthNegative)
	}
	denominator := a + b + c
	return Vector2D{
		X: (a*t.A.X + b*t.B.X + c*t.C.X) / denominator,
		Y: (a*t.A.Y + b*t.B.Y + c*t.C.Y) / denominator,
	}, nil
}

func Circumcenter(t Triangle) (Vector2D, error) {
	a := distance(t.B, t.C)
	b := distance(t.A, t.C)
	c := distance(t.A, t.B)
	area, err := HeronFormula(a, b, c)
	if err != nil {
		return Vector2D{}, err
	}
	if area == 0 {
		return Vector2D{}, errors.New(calculateFail)
	}
	D := 2 * (t.A.X*(t.B.Y-t.C.Y) + t.B.X*(t.C.Y-t.A.Y) + t.C.X*(t.A.Y-t.B.Y))
	x := ((t.A.X*t.A.X+t.A.Y*t.A.Y)*(t.B.Y-t.C.Y) +
		(t.B.X*t.B.X+t.B.Y*t.B.Y)*(t.C.Y-t.A.Y) +
		(t.C.X*t.C.X+t.C.Y*t.C.Y)*(t.A.Y-t.B.Y)) / D
	y := ((t.A.X*t.A.X+t.A.Y*t.A.Y)*(t.C.X-t.B.X) +
		(t.B.X*t.B.X+t.B.Y*t.B.Y)*(t.A.X-t.C.X) +
		(t.C.X*t.C.X+t.C.Y*t.C.Y)*(t.B.X-t.A.X)) / D
	return Vector2D{X: x, Y: y}, nil
}

func Orthocenter(t Triangle) (Vector2D, error) {
	if math.Abs(t.B.X-t.A.X) < 1e-9 {
		return Vector2D{X: t.A.X, Y: t.C.Y}, nil
	}
	if math.Abs(t.C.X-t.B.X) < 1e-9 {
		return Vector2D{X: t.B.X, Y: t.A.Y}, nil
	}
	if math.Abs(t.A.X-t.C.X) < 1e-9 {
		return Vector2D{X: t.C.X, Y: t.B.Y}, nil
	}
	slopeAB := (t.B.Y - t.A.Y) / (t.B.X - t.A.X)
	slopeBC := (t.C.Y - t.B.Y) / (t.C.X - t.B.X)
	x := (slopeAB*slopeBC*(t.A.Y-t.C.Y) + slopeBC*(t.B.X-t.A.X) - slopeAB*(t.C.X-t.B.X)) /
		(slopeBC - slopeAB)
	y := slopeAB*(x-t.A.X) + t.A.Y
	return Vector2D{X: x, Y: y}, nil
}

func HeronFormula(a, b, c float64) (float64, error) {
	if a <= 0 || b <= 0 || c <= 0 {
		return 0, errors.New(lengthNegative)
	}
	if a >= b+c || b >= a+c || c >= a+b {
		return 0, errors.New(calibrationFail)
	}
	s := (a + b + c) / 2
	areaSquared := s * (s - a) * (s - b) * (s - c)
	if areaSquared < 0 {
		return 0, errors.New(lengthNegative)
	}
	return math.Sqrt(areaSquared), nil
}

func distance(p1, p2 Vector2D) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	return math.Sqrt(dx*dx + dy*dy)
}
