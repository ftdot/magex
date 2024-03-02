package interfaces

import "image/color"

type IBgColor interface {
	SetBgColor(clr color.Color)
	GetBgColor() color.Color
}
