package camera

import (
	"github.com/ftdot/magex/components/transform"
	"github.com/ftdot/magex/utils"
	"github.com/ftdot/magex/utils/vector2"
)

type Camera struct {
	Transform transform.ITransform // Transform of the camera. Rotation doesn't implemented and will be ignored.
	ViewRange float64              // Layer limit (Z coordinate) at which the camera stops displaying the object.

	ResolutionWidth, ResolutionHeight int // Defined resolution of the screen.
	ResolutionVector                  *vector2.Vector2

	// Offset - Is a vector thattakes into account the current
	// camera coordinates and resolution size. Also takes into
	// account the camera scale.
	Offset *vector2.Vector2

	ID string // System variable with ID of the component.
}

////

func New(transform transform.ITransform, resolutionWidth, resolutionHeight int) *Camera {
	return &Camera{
		transform,
		-10,
		resolutionWidth,
		resolutionHeight,
		vector2.New(float64(resolutionWidth), float64(resolutionHeight)),
		vector2.Null.Copy(),
		utils.GenerateComponentID(),
	}
}

func (c Camera) GetID() string {
	return c.ID
}
