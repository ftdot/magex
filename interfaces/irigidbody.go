package interfaces

import "github.com/E4/box2d"

type RBType uint8

var (
	RBTypeStatic    = RBType(box2d.B2BodyType.B2_staticBody)
	RBTypeKinematic = RBType(box2d.B2BodyType.B2_kinematicBody)
	RBTypeDynamic   = RBType(box2d.B2BodyType.B2_dynamicBody)
)

////

type IRigidbody interface {
	GetType() RBType
	GetB2Body() *box2d.B2Body
	GetFixture() *box2d.B2Fixture
	UpdateTransform()
	GetID() string // Implement goi.Component
}
