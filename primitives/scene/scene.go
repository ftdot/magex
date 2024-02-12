package scene

import (
	"errors"

	"github.com/ftdot/magex/primitives/maincamera"
	"github.com/ftdot/magex/utils/interfaces"
)

////

var ErrGOMapNil = errors.New("gomap is nil, must be initialized in scene enter function")

type SceneF func(scene *Scene) error

////

type Scene struct {
	GOMap  interfaces.IGOMap
	enterF SceneF
	exitF  SceneF

	CurrentMainCamera *maincamera.MainCamera
	CurrentRigidbodies []interfaces.IRigidbody
}

func New(enterF SceneF, exitF SceneF) *Scene {
	return &Scene{
		GOMap:             nil,
		enterF:            enterF,
		exitF:             exitF,
		CurrentMainCamera: nil,
		CurrentRigidbodies: []interfaces.IRigidbody{},
	}
}

////

// Called when scene is entered by game.
func (s *Scene) Enter() error {

	if err := s.enterF(s); err != nil {
		return err
	}

	if s.GOMap == nil {
		return ErrGOMapNil
	}

	if err := s.GOMap.StartQueuedStartables(); err != nil {
		return err
	}

	return nil
}

// Called when exiting from scene.
func (s *Scene) Exit() error {

	if s.exitF != nil {
		if err := s.exitF(s); err != nil {
			return err
		}
	}

	if err := s.GOMap.UnregisterAll(); err != nil {
		return err
	}

	return nil
}

// Set ups the Enter() function handler
func (s *Scene) SetEnterF(enterF SceneF) {
	s.enterF = enterF
}

// Set ups the Exit() function handler
func (s *Scene) SetExitF(exitF SceneF) {
	s.exitF = exitF
}

////

func (s *Scene) SetMainCamera(cam *maincamera.MainCamera) error {
	s.GOMap.Unregister("MainCamera") // MainCamera may not exists
	if err := s.GOMap.Register("MainCamera", cam); err != nil {
		return err
	}
	s.CurrentMainCamera = cam

	return nil
}
