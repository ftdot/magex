package rigidbody

import (
	"github.com/E4/box2d"
	"github.com/ftdot/magex/interfaces"
	"github.com/ftdot/magex/utils"
	"github.com/ftdot/magex/utils/vector2"
)

type Rigidbody struct {
	transform interfaces.ITransform
	b2body    *box2d.B2Body
	b2fixture *box2d.B2Fixture
	bodyType  interfaces.RBType

	physSystem interfaces.PhysSystem

	id string
}

func New(
	tf interfaces.ITransform,
	ps interfaces.PhysSystem,
	bodyType interfaces.RBType,
	shape box2d.B2ShapeInterface,
) *Rigidbody {
	def := box2d.NewB2BodyDef()
	def.Type = uint8(bodyType)
	def.Position.Set(tf.GetPosition().AsTuple())

	rb := &Rigidbody{
		transform:  tf,
		b2body:     ps.NewB2Body(def),
		bodyType:   bodyType,
		physSystem: ps,
		id:         utils.GenerateComponentID(),
	}
	rb.b2fixture = rb.b2body.CreateFixture(shape, 1)
	ps.RegisterRigidbody(rb)
	return rb
}

////

func (rb *Rigidbody) GetType() interfaces.RBType {
	return rb.bodyType
}

func (rb *Rigidbody) GetTransform() interfaces.ITransform {
	return rb.transform
}

func (rb *Rigidbody) GetB2Body() *box2d.B2Body {
	return rb.b2body
}

func (rb *Rigidbody) GetFixture() *box2d.B2Fixture {
	return rb.b2fixture
}

func (rb *Rigidbody) UpdateTransform() {
	p := rb.b2body.GetPosition()
	rb.transform.SetPosition(
		vector2.NewB2(p),
	)
	rb.transform.SetRotation(
		rb.b2body.GetAngle(),
	)
}

func (rb *Rigidbody) Destroy() error {
	rb.physSystem.UnregisterRigidbody(rb)
	return nil
}

////

func (rb *Rigidbody) GetID() string {
	return rb.id
}
