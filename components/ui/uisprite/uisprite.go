package uisprite

import (
	"github.com/ftdot/magex/components/sprite"
	"github.com/ftdot/magex/components/transform"
	"github.com/ftdot/magex/game"
	"github.com/ftdot/magex/utils"
	"github.com/ftdot/magex/utils/interfaces"
	"github.com/ftdot/magex/utils/mmath"
	"github.com/ftdot/magex/utils/vector2"

	"github.com/hajimehoshi/ebiten/v2"
)

// A sprite fixed to the screen. Ignores Camera.ViewRange value.
type UISprite struct {
	Transform transform.ITransform
	Image     *ebiten.Image
	Options   interfaces.ISpriteOptions

	pivot         *vector2.Vector2
	pivotOpposite *vector2.Vector2
	imageBounds   *vector2.Vector2

	ID string // System variable with ID of the component.
}

// Creates new sprite fixed to the screen.
func New(transform transform.ITransform, image *ebiten.Image) *UISprite {
	var ix, iy float64
	if image != nil {
		ix = float64(image.Bounds().Size().X)
		iy = float64(image.Bounds().Size().Y)
	} else {
		ix = 0
		iy = 0
	}

	return &UISprite{
		transform,
		image,
		sprite.NewSpriteOptions(),
		vector2.New(-ix/2, -iy/2),
		vector2.New(ix/2, ix/2),
		vector2.New(ix, iy),
		utils.GenerateComponentID(),
	}
}

func (s *UISprite) DrawUIPriority() float64 {
	return s.Transform.GetLayer()
}

func (s *UISprite) DrawUI(game *game.GameBase, screen *ebiten.Image) {
	if !s.Options.IsVisible() {
		return
	}

	var opts ebiten.DrawImageOptions
	targetPos := s.Transform.GetPosition().Add(s.pivotOpposite.Mul(s.Transform.GetScale()))

	opts.GeoM.Scale(s.Transform.GetScale().X, s.Transform.GetScale().Y)
	opts.GeoM.Translate(s.pivot.X*s.Transform.GetScale().X, s.pivot.Y*s.Transform.GetScale().Y)
	opts.GeoM.Rotate(s.Transform.GetRotation() * mmath.RadiansMeasurement90)
	opts.GeoM.Translate(targetPos.X, targetPos.Y)

	opts.ColorScale.ScaleWithColor(s.Options.GetColor())

	screen.DrawImage(s.Image, &opts)
}

func (s *UISprite) SetTransform(t transform.ITransform) {
	s.Transform = t
}

func (s *UISprite) GetTransform() transform.ITransform {
	return s.Transform
}

func (s *UISprite) GetImage() *ebiten.Image {
	return s.Image
}

func (s *UISprite) SetImage(image *ebiten.Image) {
	s.SetPivot(vector2.New(-float64(image.Bounds().Size().X)/2, -float64(image.Bounds().Size().Y)/2))
	s.Image = image
}

func (s *UISprite) GetOptions() interfaces.ISpriteOptions {
	return s.Options
}

func (s *UISprite) SetOptions(opts interfaces.ISpriteOptions) {
	s.Options = opts
}

func (s *UISprite) GetPivot() *vector2.Vector2 {
	return s.pivot
}

func (s *UISprite) GetPivotOpposite() *vector2.Vector2 {
	return s.pivotOpposite
}

func (s *UISprite) GetPivotScaled() *vector2.Vector2 {
	return s.pivot.Mul(s.Transform.GetScale())
}

func (s *UISprite) GetPivotOppositeScaled() *vector2.Vector2 {
	return s.pivotOpposite.Mul(s.Transform.GetScale())
}

func (s *UISprite) SetPivot(pivot *vector2.Vector2) {
	s.pivot = pivot
	s.pivotOpposite = s.pivot.Reverse()
}

func (s *UISprite) GetImageBounds() *vector2.Vector2 {
	return s.imageBounds
}

func (s *UISprite) GetBoudingBox() (bbA, bbB *vector2.Vector2) {
	return nil, nil
}

func (s *UISprite) GetID() string {
	return s.ID
}
