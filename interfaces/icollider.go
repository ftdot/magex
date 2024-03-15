package interfaces

import (
	"github.com/ftdot/magex/physics/collision2d"
	"github.com/ftdot/magex/utils/ctags"
	"github.com/ftdot/magex/utils/vector2"
)

type IUICollider interface {
	SetSizeScalar(ss *vector2.Vector2)
	GetSizeScalar() *vector2.Vector2
	GetSprite() ISprite
	GetPolygon() *collision2d.Polygon
	GetCTags() *ctags.CTags
	GetPolygonAtPosition(pos *vector2.Vector2) *collision2d.Polygon
	GetID() string
}
