package gamebase

import (
	"errors"

	"github.com/ftdot/magex/interfaces"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/mobile"
)

////

var ErrNoMainScene = errors.New("main scene isn't set up")

type LayoutF func(gb *GameBase, outsideWidth, outsideHeight int) (width, height int)

////

func DefaultLayoutF(gb *GameBase, outsideWidth, outsideHeight int) (width, height int) {
	return gb.CurrentScene.GetMainCamera().GetResolutionWidth(), gb.CurrentScene.GetMainCamera().GetResolutionHeight()
}

////

type GameBase struct {
	CurrentScene interfaces.IScene // Current scene in use.

	layoutF LayoutF
}

func New() *GameBase {
	gb := &GameBase{
		nil,
		DefaultLayoutF,
	}

	return gb
}

////

func (gb *GameBase) SetScene(scene interfaces.IScene) error {

	if gb.CurrentScene != nil {
		if err := gb.CurrentScene.Exit(); err != nil {
			return err
		}
	}

	gb.CurrentScene = scene
	return scene.Enter()
}

func (gb *GameBase) GetCurrentScene() interfaces.IScene {
	return gb.CurrentScene
}

////

func (gb *GameBase) Update() error {

	// Update camera offset
	gb.CurrentScene.GetMainCamera().SetOffset(
		gb.CurrentScene.GetMainCamera().GetTransform().GetPosition().Add(gb.CurrentScene.GetMainCamera().GetResolutionVector()),
	)

	return gb.CurrentScene.GetGOMap().Update()
}

func (gb *GameBase) Draw(screen *ebiten.Image) {
	gb.CurrentScene.GetGOMap().Draw(screen)
}

func (gb *GameBase) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return gb.layoutF(gb, outsideWidth, outsideHeight)
}

func (gb *GameBase) Run(windowSizeX int, windowSizeY int, title string) error {
	if gb.CurrentScene == nil {
		return ErrNoMainScene
	}

	ebiten.SetWindowSize(windowSizeX, windowSizeY)
	ebiten.SetWindowTitle(title)
	return ebiten.RunGame(gb)
}

func (gb *GameBase) RunAsMobile() {
	mobile.SetGame(gb)
}

////
