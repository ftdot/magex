package textutil

import (
	"github.com/ftdot/magex/components/transform"
	"github.com/ftdot/magex/components/ui/uitext"
	"github.com/ftdot/magex/utils/vector2"
)

// Transform implements ITransform interface. Do not use Transform directly!
// Use interface instead of. Do not change or get values directly.
type TextTransform struct {
	*transform.Transform
	UIText *uitext.UIText // Text object to that pinned the transform.
}

////

func NewTextTransform(pos, scale *vector2.Vector2, rot, layer float64) *TextTransform {
	return &TextTransform{
		UIText:    nil,
		Transform: transform.New(pos, scale, rot, layer),
	}
}

////

func (t *TextTransform) SetUIText(uit *uitext.UIText) {
	t.UIText = uit
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
