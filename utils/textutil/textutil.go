package textutil

import (
	"github.com/ftdot/magex/utils/vector2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

func GetCenter(face font.Face, str string, pos *vector2.Vector2) *vector2.Vector2 {
	bounds := text.BoundString(face, str)
	return vector2.New(
		pos.X-float64(bounds.Min.X)-float64(bounds.Dx())/2.,
		pos.Y-float64(bounds.Min.Y)-float64(bounds.Dy())/2.,
	)
}
