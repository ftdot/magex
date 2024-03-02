package interfaces

import "github.com/ftdot/magex/utils/vector2"

type ICamera interface {
	GetTransform() ITransform
	GetViewRange() float64
	SetViewRange(viewRange float64)
	GetResolutionWidth() int
	GetResolutionHeight() int
	GetResolutionVector() *vector2.Vector2
	SetResolution(width int, height int)
	SetResolutionWidth(width int)
	SetResolutionHeight(height int)
	GetOffset() *vector2.Vector2
	SetOffset(offset *vector2.Vector2)
	GetID() string
}
