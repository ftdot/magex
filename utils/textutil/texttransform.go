package textutil

import (
	"github.com/ftdot/magex/components/transform"
	"github.com/ftdot/magex/components/ui/uitext"
	"github.com/ftdot/magex/utils"
	"github.com/ftdot/magex/utils/vector2"
)

// Transform implements ITransform interface. Do not use Transform directly!
// Use interface instead of. Do not change or get values directly.
type TextTransform struct {
	UIText   *uitext.UIText   // Text object to that pinned the transform.
	Position *vector2.Vector2 // Position. This will be aligned to the parent's position + local position, if it isn't nil.
	Scale    *vector2.Vector2 // Scale. This will be aligned to the parent's scale + local scale, if it isn't nil.
	Rotation float64          // Theta rotation in Degress.
	Layer    float64          // Also known as Z coordinate.

	LocalPosition *vector2.Vector2     // Local position. Has no any role, if parent is nil.
	LocalScale    *vector2.Vector2     // Local scale. Has no any role, if parent is nil.
	LocalRotation float64              // Local rotation. Has no any role, if parent is nil.
	LocalLayer    float64              // Local layer. Has no any role, if parent is nil.
	Parent        transform.ITransform // Parent transform (can be nil if has no parent).
	ID            string               // System variable with ID of the component.
}

////

func NewTextTransform(position, scale *vector2.Vector2, rotation, layer float64) *TextTransform {
	return &TextTransform{
		UIText:        nil,
		Position:      position,
		Scale:         scale,
		Rotation:      rotation,
		Layer:         layer,
		LocalPosition: vector2.Null.Copy(),
		LocalScale:    vector2.Null.Copy(),
		LocalRotation: 0,
		LocalLayer:    0,
		Parent:        nil,
		ID:            utils.GenerateComponentID(),
	}
}

////

func (t *TextTransform) SetUIText(uit *uitext.UIText) {
	t.UIText = uit
}

func (t *TextTransform) SetPosition(pos *vector2.Vector2) {
	if t.Parent == nil {
		t.Position = pos
		return
	}
	t.LocalPosition = t.Parent.GetPosition().Sub(pos)
}

func (t *TextTransform) SetScale(scale *vector2.Vector2) {
	if t.Parent == nil {
		t.Scale = scale
		return
	}
	t.LocalScale = t.Parent.GetScale().Div(scale)
}

func (t *TextTransform) SetRotation(rot float64) {
	if t.Parent == nil {
		t.Rotation = rot
		return
	}
	t.LocalRotation = t.Parent.GetRotation() - rot
}

func (t *TextTransform) SetLayer(layer float64) {
	if t.Parent == nil {
		t.Layer = layer
		return
	}
	t.LocalLayer = t.Parent.GetLayer() - layer
}

func (t *TextTransform) SetParent(parent transform.ITransform) {
	t.LocalPosition = t.Position
	t.LocalScale = t.Scale
	t.LocalRotation = t.Rotation
	t.LocalLayer = t.Layer
	t.Parent = parent
}

////

func (t *TextTransform) getCenterPosition(pos *vector2.Vector2) *vector2.Vector2 {
	return GetCenter(t.UIText.GetFontFace(), t.UIText.Text, pos)
}

func (t *TextTransform) GetPosition() *vector2.Vector2 {
	if t.Parent == nil {
		return t.getCenterPosition(t.Position)
	}
	return t.getCenterPosition(t.Parent.GetPosition().Add(t.LocalPosition))
}

func (t *TextTransform) GetScale() *vector2.Vector2 {
	if t.Parent == nil {
		return t.Scale
	}
	return t.Parent.GetScale().Mul(t.LocalScale)
}

func (t *TextTransform) GetRotation() float64 {
	if t.Parent == nil {
		return t.Rotation
	}
	return t.Parent.GetRotation() + t.LocalRotation
}

func (t *TextTransform) GetLayer() float64 {
	if t.Parent == nil {
		return t.Layer
	}
	return t.Parent.GetLayer() + t.LocalLayer
}

func (t *TextTransform) GetParent() transform.ITransform {
	return t.Parent
}

////

func (t *TextTransform) GetID() string {
	return t.ID
}
