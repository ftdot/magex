package shapes

import (
	"github.com/E4/box2d"
	"github.com/ftdot/magex/interfaces"
)

func NewBoxShapeFromSprite(s interfaces.ISprite) *box2d.B2PolygonShape {
	ps := box2d.NewB2PolygonShape()
	ps.SetAsBox(
		s.GetImageSize().X*s.GetTransform().GetScale().X/2,
		s.GetImageSize().Y*s.GetTransform().GetScale().Y/2,
	)
	return ps
}

func NewCircleShapeFromSprite(s interfaces.ISprite) *box2d.B2CircleShape {
	cs := box2d.NewB2CircleShape()
	cs.SetRadius(
		(s.GetImageSize().X * s.GetTransform().GetScale().X) / 2,
	)
	return cs
}
