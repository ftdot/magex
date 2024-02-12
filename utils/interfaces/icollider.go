package interfaces

import (
	"github.com/ftdot/magex/physics/collision2d"
)

type ICollider interface {
	GetSprite() ISprite
	GetPolygon() *collision2d.Polygon
}
