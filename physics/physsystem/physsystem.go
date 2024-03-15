package physsystem

import (
	"github.com/E4/box2d"
	"github.com/ftdot/magex/interfaces"
)

const defaultTimeStep float64 = 1. / 60
const defaultVelocityIterations int = 6
const defaultPositionIterations int = 2

var defaultGravity = box2d.B2Vec2{X: 0, Y: 0}

type PhysSystem struct {
	b2world      *box2d.B2World
	physSettings *interfaces.PhysSettings
	rigidbodies  map[string]interfaces.IRigidbody
}

func New() *PhysSystem {
	w := box2d.MakeB2World(defaultGravity.Clone())
	return &PhysSystem{
		b2world: &w,
		physSettings: &interfaces.PhysSettings{
			TimeStep:           defaultTimeStep,
			VelocityIterations: defaultVelocityIterations,
			PositionIterations: defaultPositionIterations,
		},
		rigidbodies: make(map[string]interfaces.IRigidbody),
	}
}

////

func (ps *PhysSystem) PhysUpdate(gb interfaces.IGameBase) error {

	ps.b2world.Step(ps.physSettings.TimeStep, ps.physSettings.VelocityIterations, ps.physSettings.PositionIterations)
	for _, rb := range ps.rigidbodies {
		rb.UpdateTransform()
	}

	return nil
}

////

func (ps *PhysSystem) NewB2Body(def *box2d.B2BodyDef) *box2d.B2Body {
	return ps.b2world.CreateBody(def)
}

func (ps *PhysSystem) RegisterRigidbody(rb interfaces.IRigidbody) {
	ps.rigidbodies[rb.GetID()] = rb
}

func (ps *PhysSystem) UnregisterRigidbody(rb interfaces.IRigidbody) {
	delete(ps.rigidbodies, rb.GetID())
}

func (ps *PhysSystem) GetB2World() *box2d.B2World {
	return ps.b2world
}
