/**
 * Author:  Nyxvectar Yan
 * Repo:    go-zju-formulas
 * Created: 07/23/2025
 */

package geometryn

type SpatialCoordinateSys struct {
	x float64
	y float64
	z float64
}

func NewSpatialPoint(x, y, z float64) SpatialCoordinateSys {
	return SpatialCoordinateSys{x, y, z}
}

func (v SpatialCoordinateSys) Add(u SpatialCoordinateSys) SpatialCoordinateSys {
	return SpatialCoordinateSys{v.x + u.x, v.y + u.y, v.z + u.z}
}

func (v SpatialCoordinateSys) Subtract(u SpatialCoordinateSys) SpatialCoordinateSys {
	return SpatialCoordinateSys{v.x - u.x, v.y - u.y, v.z - u.z}
}
