package interfaces

import (
	"github.com/ftdot/magex/physics/collision2d"
	"github.com/ftdot/magex/utils/ctags"
	"github.com/ftdot/magex/utils/vector2"
)

type ICollider interface {
	SetSizeScalar(ss *vector2.Vector2)
	GetSizeScalar() *vector2.Vector2
	GetSprite() ISprite
	GetPolygon() *collision2d.Polygon
	GetCTags() *ctags.CTags
	GetID() string
}
