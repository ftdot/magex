package interfaces

import (
	"github.com/ftdot/magex/components/transform"
	"github.com/ftdot/magex/utils/vector2"
)

type IRigidbody interface {
	GetTransform() transform.ITransform
	GetColliders() []ICollider
	GetCurrentVelocity() *vector2.Vector2
	SetCurrentVelocity(velocity *vector2.Vector2)
	GetID() string // Implement goi.Component
}
