package interfaces

import "image/color"

type ISpriteOptions interface {
	IsVisible() bool
	SetVisible(v bool)
	SetColor(c color.RGBA)
	GetColor() color.RGBA
}
