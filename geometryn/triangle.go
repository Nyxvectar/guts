/**
 * Author:  Nyxvectar Yan
 * Repo:    go-zju-formulas
 * Created: 07/23/2025
 */

// Package geometryn 提供二维几何图形（主要是三角形）的基本运算功能，包括正弦定理、余弦定理、重心、内心等几何属性的计算
package geometryn

import (
	"errors"
	"math"
)

// Vector2D 表示二维平面中的点或向量，包含X和Y两个坐标分量
type Vector2D struct {
	X float64 // X轴坐标（或向量X分量）
	Y float64 // Y轴坐标（或向量Y分量）
}

// Triangle 表示二维平面中的三角形，由三个顶点A、B、C构成
type Triangle struct {
	A Vector2D // 三角形顶点A
	B Vector2D // 三角形顶点B
	C Vector2D // 三角形顶点C
}

var (
	lengthNegative  = "三角形不存在 [length<0]"          // 边长为非正数时的错误信息
	resultNegative  = "三角形不存在 [square<0]"          // 计算结果平方为负数时的错误信息
	angleNegative   = "三角形不存在 [angle<0]"           // 角度为非正数时的错误信息
	angleOutRange   = "三角形不存在 [angleOutRange]"     // 角度超出有效范围（0到π）时的错误信息
	angleFault      = "三角形不存在 [angleTotal!=Pi]"    // 三角形内角和不等于π时的错误信息
	calibrationFail = "计算结果不一 [present!=previous]" // 验证不通过时的错误信息
	calculateFail   = "无法计算     {!}"                 // 无法计算时的错误信息
)

// LawOfSines 验证正弦定理并计算三角形外接圆半径的一半
// 正弦定理：a/sinA = b/sinB = c/sinC = 2R（R为外接圆半径）
// 参数：
//
//	a, b, c：三角形的三条边长
//	A, B, C：三条边对应的内角（弧度制）
//
// 返回值：
//
//	若验证通过，返回R（外接圆半径）；否则返回错误
//
// 错误情况：
//   - 边长为非正数（lengthNegative）
//   - 角度为非正数（angleNegative）
//   - 内角和不等于π（angleFault）
//   - 正弦定理验证不通过（calibrationFail）
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

// LawOfCosines 根据余弦定理计算三角形的第三边
// 余弦定理：c² = a² + b² - 2ab·cosC（C为a和b的夹角）
// 参数：
//
//	a, b：三角形的两条边长
//	C：a和b的夹角（弧度制，0 < C < π）
//
// 返回值：
//
//	第三边的长度；若参数无效或计算出错，返回错误
//
// 错误情况：
//   - a或b为非正数（lengthNegative）
//   - 角度C不在(0, π)范围内（angleOutRange）
//   - 计算得到的平方为负数（resultNegative）
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

// ProjectionTheorem 验证射影定理（余弦定理的特殊形式）
// 射影定理：a = b·cosC + c·cosB（a为三角形一边，B、C为另外两角）
// 参数：
//
//	a, b, c：三角形的三条边长
//	B, C：与a边相邻的两个内角（弧度制）
//
// 返回值：
//
//	若定理成立返回true；否则返回false，同时可能返回错误
//
// 错误情况：
//   - 边长为非正数（lengthNegative）
//   - 角度B或C不在(0, π)范围内，或B+C ≥ π（angleOutRange）
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

// MedianLength 计算三角形中某一边的中线长度
// 中线公式：mₐ = √[2b² + 2c² - a²]/2（mₐ为a边对应的中线）
// 参数：
//
//	a, b, c：三角形的三条边长（a为待计算中线对应的边）
//
// 返回值：
//
//	中线长度；若参数无效，返回错误
//
// 错误情况：
//   - 边长为非正数（lengthNegative）
//   - 三边无法构成三角形（calibrationFail）
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

// Centroid 计算三角形的重心（三条中线的交点）
// 重心坐标公式：( (A.X+B.X+C.X)/3, (A.Y+B.Y+C.Y)/3 )
// 参数：
//
//	t：三角形（包含三个顶点A、B、C）
//
// 返回值：
//
//	重心的二维坐标（Vector2D）
func Centroid(t Triangle) Vector2D {
	return Vector2D{
		X: (t.A.X + t.B.X + t.C.X) / 3,
		Y: (t.A.Y + t.B.Y + t.C.Y) / 3,
	}
}

// Incenter 计算三角形的内心（三条角平分线的交点，内切圆圆心）
// 内心坐标公式：( (aA.X + bB.X + cC.X)/(a+b+c), (aA.Y + bB.Y + cC.Y)/(a+b+c) )
// 其中a、b、c分别为顶点A、B、C对边的长度
// 参数：
//
//	t：三角形（包含三个顶点A、B、C）
//
// 返回值：
//
//	内心的二维坐标；若边长为非正数，返回错误
//
// 错误情况：
//   - 三角形任意边长为非正数（lengthNegative）
func Incenter(t Triangle) (Vector2D, error) {
	a := distance(t.B, t.C) // a为顶点A对边的长度
	b := distance(t.A, t.C) // b为顶点B对边的长度
	c := distance(t.A, t.B) // c为顶点C对边的长度
	if a <= 0 || b <= 0 || c <= 0 {
		return Vector2D{}, errors.New(lengthNegative)
	}
	denominator := a + b + c
	return Vector2D{
		X: (a*t.A.X + b*t.B.X + c*t.C.X) / denominator,
		Y: (a*t.A.Y + b*t.B.Y + c*t.C.Y) / denominator,
	}, nil
}

// Circumcenter 计算三角形的外心（三条边垂直平分线的交点，外接圆圆心）
// 参数：
//
//	t：三角形（包含三个顶点A、B、C）
//
// 返回值：
//
//	外心的二维坐标；若三角形面积为0（共线），返回错误
//
// 错误情况：
//   - 三角形顶点共线（面积为0），无法计算外心（calculateFail）
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

// Orthocenter 计算三角形的垂心（三条高的交点）
// 参数：
//
//	t：三角形（包含三个顶点A、B、C）
//
// 返回值：
//
//	垂心的二维坐标；处理了垂直边的特殊情况
func Orthocenter(t Triangle) (Vector2D, error) {
	// 处理垂直边的特殊情况（避免除以0）
	if math.Abs(t.B.X-t.A.X) < 1e-9 { // AB边垂直于X轴
		return Vector2D{X: t.A.X, Y: t.C.Y}, nil
	}
	if math.Abs(t.C.X-t.B.X) < 1e-9 { // BC边垂直于X轴
		return Vector2D{X: t.B.X, Y: t.A.Y}, nil
	}
	if math.Abs(t.A.X-t.C.X) < 1e-9 { // AC边垂直于X轴
		return Vector2D{X: t.C.X, Y: t.B.Y}, nil
	}

	// 一般情况计算
	slopeAB := (t.B.Y - t.A.Y) / (t.B.X - t.A.X) // AB边的斜率
	slopeBC := (t.C.Y - t.B.Y) / (t.C.X - t.B.X) // BC边的斜率
	// 垂心坐标计算
	x := (slopeAB*slopeBC*(t.A.Y-t.C.Y) + slopeBC*(t.B.X-t.A.X) - slopeAB*(t.C.X-t.B.X)) /
		(slopeBC - slopeAB)
	y := slopeAB*(x-t.A.X) + t.A.Y
	return Vector2D{X: x, Y: y}, nil
}

// HeronFormula 利用海伦公式计算三角形的面积
// 海伦公式：面积 = √[s(s-a)(s-b)(s-c)]，其中s = (a+b+c)/2
// 参数：
//
//	a, b, c：三角形的三条边长
//
// 返回值：
//
//	三角形的面积；若参数无效，返回错误
//
// 错误情况：
//   - 边长为非正数（lengthNegative）
//   - 三边无法构成三角形（calibrationFail）
//   - 计算得到的面积平方为负数（lengthNegative）
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

// distance 计算二维平面中两点间的欧氏距离
// 参数：
//
//	p1, p2：两个点的坐标（Vector2D）
//
// 返回值：
//
//	两点间的距离
func distance(p1, p2 Vector2D) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	return math.Sqrt(dx*dx + dy*dy)
}
