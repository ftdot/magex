package colliders

import (
	"image/color"

	"github.com/ftdot/magex/game"
	"github.com/ftdot/magex/physics/collision2d"
	"github.com/ftdot/magex/utils"
	"github.com/ftdot/magex/utils/ctags"
	"github.com/ftdot/magex/utils/interfaces"
	"github.com/ftdot/magex/utils/vector2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Collider that defines a circle around the sprite.
type CircleCollider struct {
	Sprite         interfaces.ISprite
	Circle         *collision2d.Circle
	PositionOffset *vector2.Vector2
	RadiusScalar   float64
	Tags           *ctags.CTags

	ID string // System variable with ID of the component.
}

func NewCircleCollider(sprite interfaces.ISprite) *CircleCollider {
	return &CircleCollider{
		Sprite:         sprite,
		Circle:         collision2d.NewCircle(vector2.Null.Copy(), -1),
		PositionOffset: vector2.Null.Copy(),
		RadiusScalar:   1,
		ID:             utils.GenerateComponentID(),
	}
}

////

func (cc *CircleCollider) SetRadiusScalar(rs float64) {
	cc.RadiusScalar = rs
}

func (cc *CircleCollider) GetRadiusScalar() float64 {
	return cc.RadiusScalar
}

func (cc *CircleCollider) SetPositionOffset(posOffset *vector2.Vector2) {
	cc.PositionOffset = posOffset
}

func (cc *CircleCollider) GetPositionOffset() *vector2.Vector2 {
	return cc.PositionOffset
}

func (cc *CircleCollider) GetPolygon() *collision2d.Polygon {
	return cc.Circle.ToPolygon()
}

func (cc *CircleCollider) GetSprite() interfaces.ISprite {
	return cc.Sprite
}

func (cc *CircleCollider) GetCTags() *ctags.CTags {
	return cc.Tags
}

////

func (cc *CircleCollider) PhysUpdate(game *game.GameBase) error {
	cc.Circle = collision2d.NewCircle(
		cc.Sprite.GetTransform().GetPosition().Add(cc.Sprite.GetPivotOppositeScaled()).Add(cc.PositionOffset),
		cc.Sprite.GetImageBounds().X*cc.Sprite.GetTransform().GetScale().X/2*cc.RadiusScalar,
	)
	return nil
}

func (cc *CircleCollider) DrawUIPriority() float64 {
	return 90000
}

// func (cc CircleCollider) DrawQueued(game *game.MagexGame, screen *ebiten.Image) {
// 	c := cc.Circle.Get()
// 	ebitenutil.DrawCircle(screen, c.Pos.X, c.Pos.Y, c.R, color.RGBA{25, 255, 25, 90})
// }

func (cc *CircleCollider) DrawUI(game *game.GameBase, screen *ebiten.Image) {

	verts := cc.Circle.ToPolygon().Points
	pos := cc.Circle.Pos.Sub(game.CurrentScene.CurrentMainCamera.Transform.GetPosition())
	for i := 0; i < len(verts); i++ {
		vert := verts[i].Add(pos)
		next := verts[0].Add(pos)

		if i < len(verts)-1 {
			next = verts[i+1].Add(pos)
		}
		ebitenutil.DrawLine(screen, vert.X, vert.Y, next.X, next.Y, color.RGBA{25, 25, 255, 255})

	}
}

////

func (bc *CircleCollider) GetID() string {
	return bc.ID
}
