package goi

import (
	"github.com/ftdot/magex/interfaces"
	"github.com/hajimehoshi/ebiten/v2"
)

type Awakable interface {
	Awake() error
}

type Destroyable interface {
	Destroy() error
}

type Startable interface {
	Start(gameBase interfaces.IGameBase) error
	StartPriority() float64
}

type PhysUpdatable interface {
	PhysUpdate(gameBase interfaces.IGameBase) error
}

type Phys2Updatable interface {
	Phys2Update(gameBase interfaces.IGameBase) error
}

type Updatable interface {
	Update(gameBase interfaces.IGameBase) error
}

type DrawableQueued interface {
	Draw(gameBase interfaces.IGameBase, screen *ebiten.Image)
	DrawPriority() float64
}

type DrawableUIQueued interface {
	DrawUI(gameBase interfaces.IGameBase, screen *ebiten.Image)
	DrawUIPriority() float64
}

type IComponent interface {
	GetID() string
}

type WithComponents interface {
	WithComponents() []IComponent
}
