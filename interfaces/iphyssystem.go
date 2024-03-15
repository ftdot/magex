package interfaces

import "github.com/E4/box2d"

type PhysSettings struct {
	TimeStep           float64
	VelocityIterations int
	PositionIterations int
}

type PhysSystem interface {
	NewB2Body(def *box2d.B2BodyDef) *box2d.B2Body
	RegisterRigidbody(rb IRigidbody)
	UnregisterRigidbody(rb IRigidbody)
	GetB2World() *box2d.B2World
}
