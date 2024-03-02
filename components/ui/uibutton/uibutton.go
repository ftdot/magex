package uibutton

import (
	"github.com/ftdot/magex/interfaces"
	"github.com/ftdot/magex/physics/collision2d"
	"github.com/ftdot/magex/utils"
	"github.com/ftdot/magex/utils/vector2"
	"github.com/hajimehoshi/ebiten/v2"
)

////

// Represents simple button that can handle hovering and clicks.
// But keep in mind, that the event fuctions can be called many
// times with different mouse buttons.
type UIButton struct {
	Sprite    interfaces.ISprite
	Colliders []interfaces.ICollider

	onHoverF   interfaces.UIButtonHoverF
	exitHoverF interfaces.UIButtonHoverF
	onClickF   interfaces.UIButtonClickF
	exitClickF interfaces.UIButtonClickF

	hovered bool
	clicked map[ebiten.MouseButton]bool

	ID string
}

func New(sprite interfaces.ISprite, colliders []interfaces.ICollider) *UIButton {
	return &UIButton{
		Sprite:    sprite,
		Colliders: colliders,
		hovered:   false,
		clicked: map[ebiten.MouseButton]bool{
			ebiten.MouseButton0: false,
			ebiten.MouseButton1: false,
			ebiten.MouseButton2: false,
			ebiten.MouseButton3: false,
			ebiten.MouseButton4: false,
		},
		ID: utils.GenerateComponentID(),
	}
}

////

func (b *UIButton) OnClick(onClickF interfaces.UIButtonClickF) {
	b.onClickF = onClickF
}

func (b *UIButton) ExitClick(exitClickF interfaces.UIButtonClickF) {
	b.exitClickF = exitClickF
}

func (b *UIButton) OnHover(onHover interfaces.UIButtonHoverF) {
	b.onHoverF = onHover
}

func (b *UIButton) ExitHover(exitHover interfaces.UIButtonHoverF) {
	b.exitHoverF = exitHover
}

////

func (b *UIButton) handleClicked(gb interfaces.IGameBase, btn ebiten.MouseButton) error {
	if ebiten.IsMouseButtonPressed(btn) {
		if !b.hovered {
			return nil
		}
		if !b.clicked[btn] {
			b.clicked[btn] = true
			if b.onClickF == nil {
				return nil
			}
			return b.onClickF(interfaces.UIButtonClickEvent{
				GameBase:    gb,
				Button:      b,
				MouseButton: btn,
			})
		}
	} else if b.clicked[btn] {
		b.clicked[btn] = false
		if b.exitClickF == nil {
			return nil
		}
		return b.exitClickF(interfaces.UIButtonClickEvent{
			GameBase:    gb,
			Button:      b,
			MouseButton: btn,
		})
	}

	return nil
}

func (b *UIButton) Update(gb interfaces.IGameBase) error {

	anyColl := false
	for _, col := range b.Colliders {
		if collision2d.PointInPolygon(vector2.NewInt(ebiten.CursorPosition()), col.GetPolygon()) {
			anyColl = true
		}
	}
	if anyColl {
		if !b.hovered {
			b.hovered = true
			if b.onHoverF == nil {
				return nil
			}
			b.onHoverF(interfaces.UIButtonHoverEvent{
				GameBase: gb,
				Button:   b,
			})
		}
	} else if b.hovered {
		b.hovered = false
		if b.exitHoverF == nil {
			return nil
		}
		b.exitHoverF(interfaces.UIButtonHoverEvent{
			GameBase: gb,
			Button:   b,
		})
	}

	if err := b.handleClicked(gb, ebiten.MouseButton0); err != nil {
		return err
	}
	if err := b.handleClicked(gb, ebiten.MouseButton1); err != nil {
		return err
	}
	if err := b.handleClicked(gb, ebiten.MouseButton2); err != nil {
		return err
	}
	if err := b.handleClicked(gb, ebiten.MouseButton3); err != nil {
		return err
	}
	if err := b.handleClicked(gb, ebiten.MouseButton4); err != nil {
		return err
	}

	return nil
}

////

func (b *UIButton) GetID() string {
	return b.ID
}
