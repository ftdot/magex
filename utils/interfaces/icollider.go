package interfaces

import (
	"github.com/ftdot/magex/physics/collision2d"
	"github.com/ftdot/magex/utils/ctags"
)

type ICollider interface {
	GetSprite() ISprite
	GetPolygon() *collision2d.Polygon
	GetCTags() *ctags.CTags
	GetID() string
}
