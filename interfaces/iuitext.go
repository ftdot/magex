package interfaces

import (
	"image/color"

	"golang.org/x/image/font"
)

////

type FontOptions struct {
	FontColor color.Color
	FontSize  float64
	DPI       float64
}

////

type IUIText interface {
	GetFontFace() font.Face
	SetColor(fontColor color.Color)
	SetDPI(dpi float64) error
	SetSize(fontSize float64) error
	GetOptions() *FontOptions
	SetOptions(fontOptions *FontOptions) error
}
