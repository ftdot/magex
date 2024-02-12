package maincamera

import (
	"github.com/ftdot/magex/components/camera"
	"github.com/ftdot/magex/components/transform"
	"github.com/ftdot/magex/utils/vector2"
)

type MainCamera struct {
	Transform transform.ITransform
	Camera    *camera.Camera
}

////

func New(resolutionWidth, resolutionHeight int) *MainCamera {
	tf := transform.New(vector2.Null.Copy(), vector2.Identity.Copy(), 0, 10)
	return &MainCamera{
		tf,
		camera.New(tf, resolutionWidth, resolutionHeight),
	}
}
