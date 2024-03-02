package interfaces

import "github.com/hajimehoshi/ebiten/v2"

type IGameBase interface {
	SetScene(s IScene) error
	GetCurrentScene() IScene
	Update() error
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int)
	Run(windowSizeX int, windowSizeY int, title string) error
	RunAsMobile()
}
