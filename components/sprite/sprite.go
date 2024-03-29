package sprite

import (
	"github.com/ftdot/magex/interfaces"
	"github.com/ftdot/magex/utils"
	"github.com/ftdot/magex/utils/mmath"
	"github.com/ftdot/magex/utils/rectutil"
	"github.com/ftdot/magex/utils/vector2"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Transform interfaces.ITransform
	Image     *ebiten.Image
	Options   interfaces.ISpriteOptions

	bbA, bbB *vector2.Vector2

	pivot         *vector2.Vector2
	pivotOpposite *vector2.Vector2
	imageBounds   *vector2.Vector2

	ID string // System variable with ID of the component
}

// Creates new sprite
func New(tf interfaces.ITransform, img *ebiten.Image) *Sprite {
	var ix, iy float64
	if img != nil {
		ix = float64(img.Bounds().Size().X)
		iy = float64(img.Bounds().Size().Y)
	} else {
		ix = 0
		iy = 0
	}

	return &Sprite{
		Transform:     tf,
		Image:         img,
		Options:       NewSpriteOptions(),
		bbA:           nil,
		bbB:           nil,
		pivot:         vector2.New(-ix/2, -iy/2),
		pivotOpposite: vector2.New(ix/2, ix/2),
		imageBounds:   vector2.New(ix, iy),
		ID:            utils.GenerateComponentID(),
	}
}

func (s Sprite) DrawPriority() float64 {
	return s.Transform.GetLayer()
}

func (s Sprite) Draw(game interfaces.IGameBase, screen *ebiten.Image) {
	if !s.Options.IsVisible() {
		return
	}
	cam := game.GetCurrentScene().GetMainCamera()
	if l := s.Transform.GetLayer(); l > cam.GetTransform().GetLayer() || l < cam.GetViewRange() {
		return
	}

	camTransform := cam.GetTransform()
	camPos := camTransform.GetPosition()
	targetPos := s.Transform.GetPosition().Sub(camPos).Add(s.GetPivotOppositeScaled())

	bA, bB := rectutil.ComputeBoundingBox(
		targetPos.X, targetPos.Y,
		s.imageBounds.X*s.Transform.GetScale().X, s.imageBounds.Y*s.Transform.GetScale().Y,
		s.Transform.GetRotation(),
	)
	bA = bA.Add(camPos)
	bB = bB.Add(camPos)
	s.bbA = bA
	s.bbB = bB

	if bA.X < cam.GetOffset().X && bB.X >= camPos.X && bB.Y >= camPos.Y && bA.Y < cam.GetOffset().Y {
		var opts ebiten.DrawImageOptions

		scale := s.Transform.GetScale()
		pivot := s.GetPivotScaled()
		opts.GeoM.Scale(scale.X, scale.Y)
		opts.GeoM.Translate(pivot.X, pivot.Y)
		opts.GeoM.Rotate(mmath.ToRadians90(s.Transform.GetRotation()))
		opts.GeoM.Translate(targetPos.X, targetPos.Y)

		opts.ColorScale.ScaleWithColor(s.Options.GetColor())

		screen.DrawImage(s.Image, &opts)
	}

}

func (s *Sprite) GetTransform() interfaces.ITransform {
	return s.Transform
}

func (s *Sprite) GetImage() *ebiten.Image {
	return s.Image
}

func (s *Sprite) SetImage(image *ebiten.Image) {
	s.SetPivot(vector2.New(-float64(image.Bounds().Size().X)/2, -float64(image.Bounds().Size().Y)/2))
	s.Image = image
}

func (s *Sprite) GetOptions() interfaces.ISpriteOptions {
	return s.Options
}

func (s *Sprite) SetOptions(opts interfaces.ISpriteOptions) {
	s.Options = opts
}

func (s *Sprite) GetPivot() *vector2.Vector2 {
	return s.pivot
}

func (s *Sprite) GetPivotOpposite() *vector2.Vector2 {
	return s.pivotOpposite
}

func (s *Sprite) GetPivotScaled() *vector2.Vector2 {
	return s.pivot.Mul(s.Transform.GetScale())
}

func (s *Sprite) GetPivotOppositeScaled() *vector2.Vector2 {
	return s.pivotOpposite.Mul(s.Transform.GetScale())
}

func (s *Sprite) SetPivot(pivot *vector2.Vector2) {
	s.pivot = pivot
	s.pivotOpposite = s.pivot.Reverse()
}

func (s *Sprite) GetImageSize() *vector2.Vector2 {
	return s.imageBounds
}

func (s *Sprite) GetBoudingBox() (bbA, bbB *vector2.Vector2) {
	return s.bbA, s.bbB
}

////

func (s *Sprite) GetID() string {
	return s.ID
}
