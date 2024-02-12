package interfaces

import "github.com/hajimehoshi/ebiten/v2"

// Not full implementation of the GOMap interface.
// There is missing SetupGame() method to avoid problems
// with an import cycle.
type IGOMap interface {
	// Tries to get the game object by given name.
	Get(name string) (i interface{}, ok bool)
	// Registers a new game object. If Awake() function implemented,
	// it will be called. This will be applied to all the components
	// of the game object.
	Register(name string, gameObject interface{}) error
	// Unregisters the game object. If Destroy() function implemented,
	// it will be called. This will be applied to all the components
	// of the game object.
	Unregister(name string) error
	// Calls Update(), PhysUpdate() functions.
	Update() error
	// Calls Draw(), DrawUI() functions, previously sorts it by
	// priority.
	Draw(screen *ebiten.Image) error
	// One-time function, that must be called from Scene.Enter()
	// only once. Starts all the queued startables, that implements
	// StartableQueued interface.
	StartQueuedStartables() error
	// Unregisters all the game objects.
	UnregisterAll() error
}
