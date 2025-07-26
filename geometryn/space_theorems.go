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

type SpatialCoordinateSys struct {
	x float64
	y float64
	z float64
}

// !! 平面方程为 ax + by + cz + d = 0
type Plane struct {
	a float64
	b float64
	c float64
	d float64
}

var (
	ErrZeroVector = errors.New("零向量没有方向")
	ErrNotCop     = errors.New("传入的两点不共面")
	ErrNotPar     = errors.New("两向量不平行")
)

func NewSpatialPoint(x, y, z float64) SpatialCoordinateSys {
	return SpatialCoordinateSys{x, y, z}
}

func (v SpatialCoordinateSys) Add(u SpatialCoordinateSys) SpatialCoordinateSys {
	return SpatialCoordinateSys{v.x + u.x, v.y + u.y, v.z + u.z}
}

func (v SpatialCoordinateSys) Subtract(u SpatialCoordinateSys) SpatialCoordinateSys {
	return SpatialCoordinateSys{v.x - u.x, v.y - u.y, v.z - u.z}
	// vectorAB := pointB.Subtract(pointA)
}

func (v SpatialCoordinateSys) Multiply(scalar float64) SpatialCoordinateSys {
	return SpatialCoordinateSys{
		v.x * scalar,
		v.y * scalar,
		v.z * scalar,
	}
}

func (v SpatialCoordinateSys) Dot(u SpatialCoordinateSys) float64 {
	return v.x*u.x + v.y*u.y + v.z*u.z
}

func (v SpatialCoordinateSys) Cross(u SpatialCoordinateSys) SpatialCoordinateSys {
	return SpatialCoordinateSys{
		v.y*u.z - v.z*u.y,
		v.z*u.x - v.x*u.z,
		v.x*u.y - v.y*u.x,
	}
}

func (v SpatialCoordinateSys) Magnitude() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v SpatialCoordinateSys) Normalize() (SpatialCoordinateSys, error) {
	var mag = v.Magnitude()
	if mag == 0 {
		return SpatialCoordinateSys{}, ErrZeroVector
	}
	return SpatialCoordinateSys{
		v.x / mag,
		v.y / mag,
		v.z / mag,
	}, nil
}

func (v SpatialCoordinateSys) IsEOrSuppleAngle(u SpatialCoordinateSys) (bool, error) {
	var vNorm, errV = v.Normalize()
	var uNorm, errU = u.Normalize()
	if errV != nil {
		return false, errV
	}
	if errU != nil {
		return false, errU
	}
	var dot = vNorm.Dot(uNorm)
	return math.Abs(math.Abs(dot)-1) < 1e-10, nil
}

func IsLineParaToPlane(line, planeNorm SpatialCoordinateSys) (bool, error) {
	if line.Magnitude() == 0 || planeNorm.Magnitude() == 0 {
		return false, ErrZeroVector
	}
	return math.Abs(line.Dot(planeNorm)) < 1e-10, nil
}

func NewPlane(pa, pb, pc SpatialCoordinateSys) (Plane, error) {
	var vAB = pb.Subtract(pa)
	var vAC = pc.Subtract(pa)
	var normal = vAB.Cross(vAC)
	if normal.Magnitude() == 0 {
		return Plane{}, ErrNotCop
	}
	var a, b, c = normal.x, normal.y, normal.z
	var d = -(a*pa.x + b*pa.y + c*pa.z)
	// 平面点法式方程的推导原理：
	// 平面的法向量(normal)与平面内任意向量垂直
	// 已知平面上一点P0(x0,y0,z0)和法向量(A,B,C)
	// 对于平面上任意点P(x,y,z)，向量PP0与法向量
	// 点积为0即: A(x-x0) + B(y-y0) + C(z-z0)= 0
	// 展开后得到一般式: Ax + By + Cz + D = 0.其
	// 中D=-(Ax0+By0+Cz0)用点pa(x0,y0,z0)计算得D
	return Plane{
		a,
		b,
		c,
		d,
	}, nil
}

func (p Plane) Normal() SpatialCoordinateSys {
	return SpatialCoordinateSys{p.a, p.b, p.c}
	// 到这里仍然是对平面一般式的运用，
	// 平面一般方程的定义中A,B,C 三个值
	// 本身就定义了其法向量，所以在此我
	// 们可以高效的求出她。
}

func ArePlanesParallel(p1, p2 Plane) (bool, error) {
	n1 := p1.Normal()
	n2 := p2.Normal()

	if n1.Magnitude() == 0 || n2.Magnitude() == 0 {
		return false, ErrZeroVector
	}
	cross := n1.Cross(n2)
	// 利用的是两个平面的法向量平行来判定
	return cross.Magnitude() < 1e-10, nil
}

func GetPlanesIntersLine(p1, p2, intersectingPlane Plane) (SpatialCoordinateSys, SpatialCoordinateSys, error) {
	areParallel, err := ArePlanesParallel(p1, p2)
	if err != nil {
		return SpatialCoordinateSys{}, SpatialCoordinateSys{}, err
	}
	if !areParallel {
		return SpatialCoordinateSys{}, SpatialCoordinateSys{}, ErrNotPar
	}
	n1 := p1.Normal()
	n2 := intersectingPlane.Normal()
	lineDir1 := n1.Cross(n2)
	n3 := p2.Normal()
	lineDir2 := n3.Cross(n2)
	return lineDir1, lineDir2, nil
}

func GetLinePlaneIntersLine(lineDir SpatialCoordinateSys, plane Plane) (SpatialCoordinateSys, error) {
	normal := plane.Normal()
	if lineDir.Magnitude() == 0 || normal.Magnitude() == 0 {
		return SpatialCoordinateSys{}, ErrZeroVector
	}
	isParallel, err := IsLineParaToPlane(lineDir, normal)
	if err != nil {
		return SpatialCoordinateSys{}, err
	}
	if !isParallel {
		return SpatialCoordinateSys{}, ErrNotPar
	}
	return lineDir.Cross(normal), nil
}

func IsLinePerToPlane(lineDir SpatialCoordinateSys, plane Plane) (bool, error) {
	if lineDir.Magnitude() == 0 {
		return false, ErrZeroVector
	}
	normal := plane.Normal()
	if normal.Magnitude() == 0 {
		return false, ErrZeroVector
	}
	cross := lineDir.Cross(normal)
	return cross.Magnitude() < 10, nil
}
