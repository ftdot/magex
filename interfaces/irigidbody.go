package interfaces

import (
	"github.com/ftdot/magex/utils/ctags"
	"github.com/ftdot/magex/utils/vector2"
)

////

type CollisionEvent struct {
	GameBase IGameBase
	OtherRB  IRigidbody
	Tags     *ctags.CTags
}

type CollFunc func(coll CollisionEvent) error

////

type IRigidbody interface {
	GetTransform() ITransform
	GetColliders() []ICollider
	SetColliders(colliders []ICollider)
	GetMass() float64
	SetMass(mass float64)
	GetCurrentVelocity() *vector2.Vector2
	SetCurrentVelocity(velocity *vector2.Vector2)
	OnCollision(ocf CollFunc)
	ExitCollision(ecf CollFunc)
	Activate(gb IGameBase)
	Deactivate(gb IGameBase)
	GetID() string // Implement goi.Component
}
