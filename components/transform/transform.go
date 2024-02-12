package transform

import (
	"github.com/ftdot/magex/utils"
	"github.com/ftdot/magex/utils/vector2"
)

// Transform implements ITransform interface. Do not use Transform directly!
// Use interface instead of. Do not change or get values directly.
//
// This component does not need to be added to the GetComponents() function!
type Transform struct {
	Position *vector2.Vector2 // Position. This will be aligned to the parent's position + local position, if it isn't nil
	Scale    *vector2.Vector2 // Scale. This will be aligned to the parent's scale + local scale, if it isn't nil
	Rotation float64          // Theta rotation in Degress
	Layer    float64          // Also known as Z coordinate

	LocalPosition *vector2.Vector2 // Local position. Has no any role, if parent is nil
	LocalScale    *vector2.Vector2 // Local scale. Has no any role, if parent is nil
	LocalRotation float64          // Local rotation. Has no any role, if parent is nil
	LocalLayer    float64          // Local layer. Has no any role, if parent is nil
	Parent        ITransform       // Parent transform (can be nil if has no parent)
	ID            string           // System variable with ID of the component
}

////

func New(position, scale *vector2.Vector2, rotation, layer float64) *Transform {
	return &Transform{
		position,
		scale,
		rotation,
		layer,
		vector2.Null.Copy(),
		vector2.Null.Copy(),
		0,
		0,
		nil,
		utils.GenerateComponentID(),
	}
}

////

func (t *Transform) SetPosition(pos *vector2.Vector2) {
	if (t.Parent == nil) {
		t.Position = pos
		return
	}
	t.LocalPosition = t.Parent.GetPosition().Sub(pos)
}

func (t *Transform) SetScale(scale *vector2.Vector2) {
	if (t.Parent == nil) {
		t.Scale = scale
		return
	}
	t.LocalScale = t.Parent.GetScale().Div(scale)
}

func (t *Transform) SetRotation(rot float64) {
	if (t.Parent == nil) {
		t.Rotation = rot
		return
	}
	t.LocalRotation = t.Parent.GetRotation() - rot
}

func (t *Transform) SetLayer(layer float64) {
	if (t.Parent == nil) {
		t.Layer = layer
		return
	}
	t.LocalLayer = t.Parent.GetLayer() - layer
}

func (t *Transform) SetParent(parent ITransform) {
	t.LocalPosition = t.Position
	t.LocalScale = t.Scale
	t.LocalRotation = t.Rotation
	t.LocalLayer = t.Layer
	t.Parent = parent
}

////

func (t *Transform) GetPosition() *vector2.Vector2 {
	if t.Parent == nil {
		return t.Position
	}
	return t.Parent.GetPosition().Add(t.LocalPosition)
}

func (t *Transform) GetScale() *vector2.Vector2 {
	if t.Parent == nil {
		return t.Scale
	}
	return t.Parent.GetScale().Mul(t.LocalScale)
}

func (t *Transform) GetRotation() float64 {
	if t.Parent == nil {
		return t.Rotation
	}
	return t.Parent.GetRotation() + t.LocalRotation
}

func (t *Transform) GetLayer() float64 {
	if t.Parent == nil {
		return t.Layer
	}
	return t.Parent.GetLayer() + t.LocalLayer
}

func (t *Transform) GetParent() ITransform {
	return t.Parent
}

////

func (t Transform) GetID() string {
	return t.ID
}
