package interfaces

import (
	"github.com/ftdot/magex/utils/vector2"

	"github.com/hajimehoshi/ebiten/v2"
)

type ISprite interface {
	GetTransform() ITransform
	GetImage() *ebiten.Image
	SetImage(image *ebiten.Image)
	GetOptions() ISpriteOptions
	SetOptions(opts ISpriteOptions)
	GetPivot() *vector2.Vector2
	GetPivotOpposite() *vector2.Vector2
	GetPivotScaled() *vector2.Vector2
	GetPivotOppositeScaled() *vector2.Vector2
	SetPivot(pivot *vector2.Vector2)
	GetImageSize() *vector2.Vector2
	GetBoudingBox() (bbA, bbB *vector2.Vector2)
	GetID() string
}
