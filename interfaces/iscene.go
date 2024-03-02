package interfaces

type SceneF func(scene IScene) error

type IScene interface {
	Enter() error
	Exit() error
	SetEnterF(enterF SceneF)
	SetExitF(exitF SceneF)
	SetMainCamera(cam ICamera) error
	GetMainCamera() ICamera
	GetGOMap() IGOMap

	SetCurrentRigidbodies([]IRigidbody)
	GetCurrentRigidbodies() []IRigidbody
}
