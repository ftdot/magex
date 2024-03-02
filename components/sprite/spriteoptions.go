package sprite

import "image/color"

type SpriteOptions struct {
	Visible bool
	Color   color.Color
}

func NewSpriteOptions() *SpriteOptions {
	return &SpriteOptions{
		Visible: true,
		Color:   color.RGBA{255, 255, 255, 255},
	}
}

////

func (so *SpriteOptions) IsVisible() bool {
	return so.Visible
}

func (so *SpriteOptions) GetColor() color.Color {
	return so.Color
}

func (so *SpriteOptions) SetVisible(v bool) {
	so.Visible = v
}

func (so *SpriteOptions) SetColor(c color.Color) {
	so.Color = c
}
