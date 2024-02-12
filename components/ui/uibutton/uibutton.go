package uibutton

import (
	"github.com/ftdot/magex/components/ui/uisprite"
	"github.com/ftdot/magex/game"
	"github.com/ftdot/magex/physics/collision2d"
	"github.com/ftdot/magex/utils"
	"github.com/ftdot/magex/utils/interfaces"
	"github.com/ftdot/magex/utils/vector2"
	"github.com/hajimehoshi/ebiten/v2"
)

////

type UIButtonClickEvent struct {
	GameBase    *game.GameBase
	Button      *UIButton
	MouseButton ebiten.MouseButton
}

////

type HoverF func(gb *game.GameBase, uiButton *UIButton) error
type ClickF func(e UIButtonClickEvent) error

////

// Represents simple button that can handle hovering and clicks.
// But keep in mind, that the event fuctions can be called many
// times with different mouse buttons.
type UIButton struct {
	UISprite  *uisprite.UISprite
	Colliders []interfaces.ICollider

	onHoverF   HoverF
	exitHoverF HoverF
	onClickF   ClickF
	exitClickF ClickF

	hovered bool
	clicked map[ebiten.MouseButton]bool

	ID string
}

func New(uiSprite *uisprite.UISprite, cols []interfaces.ICollider) *UIButton {
	return &UIButton{
		UISprite:  uiSprite,
		Colliders: cols,
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

func (b *UIButton) OnClick(onClickF ClickF) {
	b.onClickF = onClickF
}

func (b *UIButton) ExitClick(exitClickF ClickF) {
	b.exitClickF = exitClickF
}

func (b *UIButton) OnHover(onHover HoverF) {
	b.onHoverF = onHover
}

func (b *UIButton) ExitHover(exitHover HoverF) {
	b.exitHoverF = exitHover
}

////

func (b *UIButton) handleClicked(gb *game.GameBase, btn ebiten.MouseButton) error {
	if ebiten.IsMouseButtonPressed(btn) {
		if !b.hovered {
			return nil
		}
		if !b.clicked[btn] {
			b.clicked[btn] = true
			if b.onClickF == nil {
				return nil
			}
			return b.onClickF(UIButtonClickEvent{GameBase: gb, Button: b, MouseButton: btn})
		}
	} else if b.clicked[btn] {
		b.clicked[btn] = false
		if b.exitClickF == nil {
			return nil
		}
		return b.exitClickF(UIButtonClickEvent{GameBase: gb, Button: b, MouseButton: btn})
	}

	return nil
}

func (b *UIButton) Update(gb *game.GameBase) error {

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
			b.onHoverF(gb, b)
		}
	} else if b.hovered {
		b.hovered = false
		if b.exitHoverF == nil {
			return nil
		}
		b.exitHoverF(gb, b)
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
