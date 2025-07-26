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

// 平面方程为ax + by + cz + d = 0
type Plane struct {
	a float64
	b float64
	c float64
	d float64
}

var (
	ErrZeroVector = "零向量没有方向"
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
		return SpatialCoordinateSys{}, errors.New(ErrZeroVector)
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
		return false, errors.New(ErrZeroVector)
	}
	return math.Abs(line.Dot(planeNorm)) < 1e-10, nil
}
