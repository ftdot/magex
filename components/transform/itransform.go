package transform

import "github.com/ftdot/magex/utils/vector2"

type ITransform interface {
	SetPosition(pos *vector2.Vector2)
	SetScale(scale *vector2.Vector2)
	SetRotation(rot float64)
	SetLayer(layer float64)
	SetParent(parent ITransform)

	GetPosition() *vector2.Vector2
	GetScale() *vector2.Vector2
	GetRotation() float64
	GetLayer() float64
	GetParent() ITransform
}
