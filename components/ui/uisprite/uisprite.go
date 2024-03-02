package uisprite

import (
	"github.com/ftdot/magex/components/sprite"
	"github.com/ftdot/magex/interfaces"
	"github.com/ftdot/magex/utils/mmath"

	"github.com/hajimehoshi/ebiten/v2"
)

// A sprite fixed to the screen. Ignores Camera.ViewRange value.
type UISprite struct {
	*sprite.Sprite
}

// Creates new sprite fixed to the screen.
func New(tf interfaces.ITransform, img *ebiten.Image) *UISprite {
	return &UISprite{
		Sprite: sprite.New(tf, img),
	}
}

func (s *UISprite) DrawUI(gb interfaces.IGameBase, screen *ebiten.Image) {
	if !s.Options.IsVisible() {
		return
	}

	var opts ebiten.DrawImageOptions
	targetPos := s.Transform.GetPosition().Add(s.GetPivotOppositeScaled())

	opts.GeoM.Scale(s.Transform.GetScale().X, s.Transform.GetScale().Y)
	opts.GeoM.Translate(s.GetPivotScaled().X, s.GetPivotScaled().Y)
	opts.GeoM.Rotate(s.Transform.GetRotation() * mmath.RadiansMeasurement90)
	opts.GeoM.Translate(targetPos.X, targetPos.Y)

	opts.ColorScale.ScaleWithColor(s.Options.GetColor())

	screen.DrawImage(s.Image, &opts)
}
