package game

import (
	"errors"

	"github.com/ftdot/magex/primitives/scene"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/mobile"
)

////

var ErrNoMainScene = errors.New("main scene isn't set up")

////

type GameBase struct {
	CurrentScene *scene.Scene // Current scene in use.
}

func New() *GameBase {
	g := &GameBase{
		nil,
	}

	return g
}

////

func (g *GameBase) SetScene(scene *scene.Scene) error {

	if g.CurrentScene != nil {
		if err := g.CurrentScene.Exit(); err != nil {
			return err
		}
	}

	g.CurrentScene = scene
	return scene.Enter()
}

////

func (g *GameBase) Update() error {

	// Update camera offset
	g.CurrentScene.CurrentMainCamera.Camera.Offset = g.CurrentScene.CurrentMainCamera.Transform.GetPosition().Add(g.CurrentScene.CurrentMainCamera.Camera.ResolutionVector)

	return g.CurrentScene.GOMap.Update()
}

func (g *GameBase) Draw(screen *ebiten.Image) {
	g.CurrentScene.GOMap.Draw(screen)
}

func (g *GameBase) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.CurrentScene.CurrentMainCamera.Camera.ResolutionWidth, g.CurrentScene.CurrentMainCamera.Camera.ResolutionHeight
}

func (g *GameBase) Run(windowSizeX int, windowSizeY int, title string) error {
	if g.CurrentScene == nil {
		return ErrNoMainScene
	}

	ebiten.SetWindowSize(windowSizeX, windowSizeY)
	ebiten.SetWindowTitle(title)
	return ebiten.RunGame(g)
}

func (g *GameBase) RunAsMobile() {
	mobile.SetGame(g)
}

////
