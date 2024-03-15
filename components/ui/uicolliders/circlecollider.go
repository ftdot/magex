package uicolliders

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
type UICircleCollider struct {
	Sprite         interfaces.ISprite
	CirclePolygon  *collision2d.Polygon
	PositionOffset *vector2.Vector2
	SizeScalar     *vector2.Vector2
	Tags           *ctags.CTags

	ID string // System variable with ID of the component.
}

func NewUICircleCollider(sprite interfaces.ISprite) *UICircleCollider {
	return &UICircleCollider{
		Sprite:         sprite,
		CirclePolygon:  collision2d.NewCircle(vector2.Null.Copy(), -1, vector2.Null.Copy()).ToPolygon(),
		PositionOffset: vector2.Null.Copy(),
		SizeScalar:     vector2.Identity.Copy(),
		Tags:           ctags.New(),
		ID:             utils.GenerateComponentID(),
	}
}

////

func (cc *UICircleCollider) SetSizeScalar(ss *vector2.Vector2) {
	cc.SizeScalar = ss
}

func (cc *UICircleCollider) GetSizeScalar() *vector2.Vector2 {
	return cc.SizeScalar
}

func (cc *UICircleCollider) SetPositionOffset(posOffset *vector2.Vector2) {
	cc.PositionOffset = posOffset
}

func (cc *UICircleCollider) GetPositionOffset() *vector2.Vector2 {
	return cc.PositionOffset
}

func (cc *UICircleCollider) GetPolygon() *collision2d.Polygon {
	return cc.CirclePolygon
}

func (cc *UICircleCollider) GetSprite() interfaces.ISprite {
	return cc.Sprite
}

func (cc *UICircleCollider) GetCTags() *ctags.CTags {
	return cc.Tags
}

func (cc *UICircleCollider) GetPolygonAtPosition(pos *vector2.Vector2) *collision2d.Polygon {
	return collision2d.NewCircle(
		pos.Add(cc.Sprite.GetPivotOppositeScaled()).Add(cc.PositionOffset),
		cc.Sprite.GetImageSize().X*cc.Sprite.GetTransform().GetScale().X/2,
		cc.SizeScalar,
	).ToPolygon()
}

////

func (cc *UICircleCollider) PhysUpdate(gb interfaces.IGameBase) error {
	cc.CirclePolygon = cc.GetPolygonAtPosition(cc.Sprite.GetTransform().GetPosition())
	return nil
}

func (cc *UICircleCollider) DrawUIPriority() float64 {
	return 90000
}

// func (cc CircleCollider) DrawQueued(game *game.MagexGame, screen *ebiten.Image) {
// 	c := cc.Circle.Get()
// 	ebitenutil.DrawCircle(screen, c.Pos.X, c.Pos.Y, c.R, color.RGBA{25, 255, 25, 90})
// }

func (cc *UICircleCollider) DrawUI(gb interfaces.IGameBase, screen *ebiten.Image) {

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

func (bc *UICircleCollider) GetID() string {
	return bc.ID
}
