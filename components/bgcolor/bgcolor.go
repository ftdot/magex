package bgcolor

import (
	"image/color"
	"math"

	"github.com/ftdot/magex/interfaces"
	"github.com/ftdot/magex/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

var drawPriority float64 = math.Inf(-1)

type BgColor struct {
	BgColor color.Color

	ID string // System variable with ID of the component.
}

func New(color color.Color) *BgColor {
	return &BgColor{
		BgColor: color,
		ID:      utils.GenerateComponentID(),
	}
}

////

func (bc *BgColor) DrawPriority() float64 {
	return drawPriority
}

func (bc *BgColor) Draw(gb interfaces.IGameBase, screen *ebiten.Image) {
	screen.Fill(bc.BgColor)
}

////

func (bc *BgColor) SetBgColor(clr color.Color) {
	bc.BgColor = clr
}

func (bc *BgColor) GetBgColor() color.Color {
	return bc.BgColor
}

////

func (s *BgColor) GetID() string {
	return s.ID
}
