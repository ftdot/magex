package goi

import (
	"github.com/ftdot/magex/game"

	"github.com/hajimehoshi/ebiten/v2"
)

type Awakable interface {
	Awake() error
}

type Destroyable interface {
	Destroy() error
}

type Startable interface {
	Start(gameBase *game.GameBase) error
	StartPriority() float64
}

type PhysUpdatable interface {
	PhysUpdate(gameBase *game.GameBase) error
}

type Phys2Updatable interface {
	Phys2Update(gameBase *game.GameBase) error
}

type Updatable interface {
	Update(gameBase *game.GameBase) error
}

type DrawableQueued interface {
	Draw(gameBase *game.GameBase, screen *ebiten.Image)
	DrawPriority() float64
}

type DrawableUIQueued interface {
	DrawUI(gameBase *game.GameBase, screen *ebiten.Image)
	DrawUIPriority() float64
}

type Component interface {
	GetID() string
}

type WithComponents interface {
	GetComponents() []Component
}
