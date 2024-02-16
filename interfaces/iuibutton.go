package interfaces

import "github.com/hajimehoshi/ebiten/v2"

////

type UIButtonHoverEvent struct {
	GameBase IGameBase
	Button IUIButton
}

type UIButtonClickEvent struct {
	GameBase    IGameBase
	Button      IUIButton
	MouseButton ebiten.MouseButton
}

////

type UIButtonClickF func(e UIButtonClickEvent) error
type UIButtonHoverF func(e UIButtonHoverEvent) error

////

type IUIButton interface {
	OnClick(onClickF UIButtonClickF)
	ExitClick(exitClickF UIButtonClickF)
	OnHover(onHover UIButtonHoverF)
	ExitHover(exitHover UIButtonHoverF)
	GetID() string
}
