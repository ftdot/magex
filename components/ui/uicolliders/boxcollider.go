package uicolliders

import (
	"image/color"
	"math"

	"github.com/ftdot/magex/interfaces"
	"github.com/ftdot/magex/physics/collision2d"
	"github.com/ftdot/magex/utils"
	"github.com/ftdot/magex/utils/ctags"
	"github.com/ftdot/magex/utils/mmath"
	"github.com/ftdot/magex/utils/vector2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Collider that defines a box around the sprite.
type UIBoxCollider struct {
	Sprite         interfaces.ISprite
	BoxPolygon     *collision2d.Polygon
	PositionOffset *vector2.Vector2
	SizeScalar     *vector2.Vector2
	Tags           *ctags.CTags

	ID string // System variable with ID of the component.
}

func NewUIBoxCollider(sprite interfaces.ISprite) *UIBoxCollider {
	return &UIBoxCollider{
		Sprite:         sprite,
		BoxPolygon:     collision2d.NewBox(vector2.Null.Copy(), -1, -1).ToPolygon(),
		PositionOffset: vector2.Null.Copy(),
		SizeScalar:     vector2.Identity.Copy(),
		Tags:           ctags.New(),
		ID:             utils.GenerateComponentID(),
	}
}

////

func (bc *UIBoxCollider) SetSizeScalar(ss *vector2.Vector2) {
	bc.SizeScalar = ss
}

func (bc *UIBoxCollider) GetSizeScalar() *vector2.Vector2 {
	return bc.SizeScalar
}

func (bc *UIBoxCollider) SetPositionOffset(posOffset *vector2.Vector2) {
	bc.PositionOffset = posOffset
}

func (bc *UIBoxCollider) GetPositionOffset() *vector2.Vector2 {
	return bc.PositionOffset
}

func (bc *UIBoxCollider) GetPolygon() *collision2d.Polygon {
	return bc.BoxPolygon
}

func (bc *UIBoxCollider) GetSprite() interfaces.ISprite {
	return bc.Sprite
}

func (bc *UIBoxCollider) GetCTags() *ctags.CTags {
	return bc.Tags
}

func (bc *UIBoxCollider) GetPolygonAtPosition(pos *vector2.Vector2) *collision2d.Polygon {
	bounds := bc.Sprite.GetImageSize().Mul(bc.Sprite.GetTransform().GetScale()).Mul(bc.SizeScalar)
	pivot := bc.Sprite.GetPivotScaled()
	pivotOpposite := bc.Sprite.GetPivotOpposite().Mul(bc.Sprite.GetTransform().GetScale())

	boxPol := collision2d.NewBox(pos.Add(bc.PositionOffset), bounds.X, bounds.Y).ToPolygon()
	boxPol.Translate(pivot.X, pivot.Y)
	boxPol.Rotate(math.Mod(bc.Sprite.GetTransform().GetRotation(), 360) * mmath.RadiansMeasurement90)
	boxPol.Translate(pivotOpposite.X, pivotOpposite.Y)

	return boxPol
}

////

func (bc *UIBoxCollider) PhysUpdate(gb interfaces.IGameBase) error {
	bc.BoxPolygon = bc.GetPolygonAtPosition(bc.Sprite.GetTransform().GetPosition())

	return nil
}

func (bc *UIBoxCollider) DrawUIPriority() float64 {
	return 90000
}

func (bc *UIBoxCollider) DrawUI(gb interfaces.IGameBase, screen *ebiten.Image) {

	verts := bc.BoxPolygon.Points
	pos := bc.BoxPolygon.Pos
	for i := 0; i < len(verts); i++ {
		vert := verts[i].Add(pos)
		next := verts[0].Add(pos)

		if i < len(verts)-1 {
			next = verts[i+1].Add(pos)
		}
		ebitenutil.DrawLine(screen, vert.X, vert.Y, next.X, next.Y, color.RGBA{25, 255, 25, 255})

	}
}

////

func (bc *UIBoxCollider) GetID() string {
	return bc.ID
}
