package scene

import (
	"github.com/ftdot/magex/interfaces"
	"github.com/ftdot/magex/physics/physsystem"
	"github.com/ftdot/magex/utils/gomap"
)

////

type Scene struct {
	gomap      interfaces.IGOMap
	physSystem interfaces.PhysSystem
	enterF     interfaces.SceneF
	exitF      interfaces.SceneF

	CurrentMainCamera  interfaces.ICamera
	CurrentRigidbodies []interfaces.IRigidbody
}

func New(enterF interfaces.SceneF, exitF interfaces.SceneF) *Scene {
	return &Scene{
		gomap:              gomap.New(),
		physSystem:         physsystem.New(),
		enterF:             enterF,
		exitF:              exitF,
		CurrentMainCamera:  nil,
		CurrentRigidbodies: []interfaces.IRigidbody{},
	}
}

////

// Called when scene is entered by game.
func (s *Scene) Enter() error {

	s.gomap = gomap.New()
	if err := s.gomap.Register("Physics System", s.physSystem); err != nil {
		return err
	}

	if err := s.enterF(s); err != nil {
		return err
	}

	if err := s.gomap.StartQueuedStartables(); err != nil {
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

	if err := s.gomap.UnregisterAll(); err != nil {
		return err
	}

	return nil
}

// Set ups the Enter() function handler
func (s *Scene) SetEnterF(enterF interfaces.SceneF) {
	s.enterF = enterF
}

// Set ups the Exit() function handler
func (s *Scene) SetExitF(exitF interfaces.SceneF) {
	s.exitF = exitF
}

////

func (s *Scene) SetMainCamera(cam interfaces.ICamera) error {
	s.gomap.Unregister("MainCamera") // MainCamera may not exists
	if err := s.gomap.Register("MainCamera", cam); err != nil {
		return err
	}
	s.CurrentMainCamera = cam

	return nil
}

////

func (s *Scene) GetMainCamera() interfaces.ICamera {
	return s.CurrentMainCamera
}

func (s *Scene) GetGOMap() interfaces.IGOMap {
	return s.gomap
}

func (s *Scene) GetPhysSystem() interfaces.PhysSystem {
	return s.physSystem
}
