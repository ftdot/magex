package colliders

import (
	"image/color"

	"github.com/ftdot/magex/interfaces"
	"github.com/ftdot/magex/physics/collision2d"
	"github.com/ftdot/magex/utils"
	"github.com/ftdot/magex/utils/ctags"
	"github.com/ftdot/magex/utils/vector2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Collider that defines a circle around the sprite.
type CircleCollider struct {
	Sprite         interfaces.ISprite
	CirclePolygon  *collision2d.Polygon
	PositionOffset *vector2.Vector2
	SizeScalar     *vector2.Vector2
	Tags           *ctags.CTags

	ID string // System variable with ID of the component.
}

func NewCircleCollider(sprite interfaces.ISprite) *CircleCollider {
	return &CircleCollider{
		Sprite:         sprite,
		CirclePolygon:  collision2d.NewCircle(vector2.Null.Copy(), -1, vector2.Null.Copy()).ToPolygon(),
		PositionOffset: vector2.Null.Copy(),
		SizeScalar:     vector2.Identity.Copy(),
		Tags:           ctags.New(),
		ID:             utils.GenerateComponentID(),
	}
}

////

func (cc *CircleCollider) SetSizeScalar(ss *vector2.Vector2) {
	cc.SizeScalar = ss
}

func (cc *CircleCollider) GetSizeScalar() *vector2.Vector2 {
	return cc.SizeScalar
}

func (cc *CircleCollider) SetPositionOffset(posOffset *vector2.Vector2) {
	cc.PositionOffset = posOffset
}

func (cc *CircleCollider) GetPositionOffset() *vector2.Vector2 {
	return cc.PositionOffset
}

func (cc *CircleCollider) GetPolygon() *collision2d.Polygon {
	return cc.CirclePolygon
}

func (cc *CircleCollider) GetSprite() interfaces.ISprite {
	return cc.Sprite
}

func (cc *CircleCollider) GetCTags() *ctags.CTags {
	return cc.Tags
}

////

func (cc *CircleCollider) PhysUpdate(gb interfaces.IGameBase) error {
	cc.CirclePolygon = collision2d.NewCircle(
		cc.Sprite.GetTransform().GetPosition().Add(cc.Sprite.GetPivotOppositeScaled()).Add(cc.PositionOffset),
		cc.Sprite.GetImageSize().X*cc.Sprite.GetTransform().GetScale().X/2,
		cc.SizeScalar,
	).ToPolygon()
	return nil
}

func (cc *CircleCollider) DrawUIPriority() float64 {
	return 90000
}

// func (cc CircleCollider) DrawQueued(game *game.MagexGame, screen *ebiten.Image) {
// 	c := cc.Circle.Get()
// 	ebitenutil.DrawCircle(screen, c.Pos.X, c.Pos.Y, c.R, color.RGBA{25, 255, 25, 90})
// }

func (cc *CircleCollider) DrawUI(gb interfaces.IGameBase, screen *ebiten.Image) {

	verts := cc.CirclePolygon.Points
	pos := cc.CirclePolygon.Pos.Sub(gb.GetCurrentScene().GetMainCamera().GetTransform().GetPosition())
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
