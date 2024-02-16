package camera

import (
	"github.com/ftdot/magex/interfaces"
	"github.com/ftdot/magex/utils"
	"github.com/ftdot/magex/utils/vector2"
)

type Camera struct {
	Transform interfaces.ITransform // Transform of the camera. Rotation doesn't implemented and will be ignored.
	ViewRange float64               // Layer limit (Z coordinate) at which the camera stops displaying the object.

	ResolutionWidth, ResolutionHeight int // Defined resolution of the screen.
	ResolutionVector                  *vector2.Vector2

	// Offset - Is a vector thattakes into account the current
	// camera coordinates and resolution size. Also takes into
	// account the camera scale.
	Offset *vector2.Vector2

	ID string // System variable with ID of the component.
}

////

func New(tf interfaces.ITransform, resolutionWidth, resolutionHeight int) *Camera {
	return &Camera{
		Transform:        tf,
		ViewRange:        -10,
		ResolutionWidth:  resolutionWidth,
		ResolutionHeight: resolutionHeight,
		ResolutionVector: vector2.New(float64(resolutionWidth), float64(resolutionHeight)),
		Offset:           vector2.Null.Copy(),
		ID:               utils.GenerateComponentID(),
	}
}

////

func (c *Camera) GetTransform() interfaces.ITransform {
	return c.Transform
}

func (c *Camera) GetViewRange() float64 {
	return c.ViewRange
}

func (c *Camera) SetViewRange(viewRange float64) {
	c.ViewRange = viewRange
}

func (c *Camera) GetResolutionWidth() int {
	return c.ResolutionWidth
}

func (c *Camera) GetResolutionHeight() int {
	return c.ResolutionHeight
}

func (c *Camera) GetResolutionVector() *vector2.Vector2 {
	return c.ResolutionVector
}

func (c *Camera) SetResolution(width int, height int) {
	c.ResolutionWidth = width
	c.ResolutionHeight = height
	c.ResolutionVector = vector2.New(float64(width), float64(height))
}

func (c *Camera) SetResolutionWidth(width int) {
	c.ResolutionWidth = width
}

func (c *Camera) SetResolutionHeight(height int) {
	c.ResolutionHeight = height
}

func (c *Camera) GetOffset() *vector2.Vector2 {
	return c.Offset
}

func (c *Camera) SetOffset(offset *vector2.Vector2) {
	c.Offset = offset
}

////

func (c Camera) GetID() string {
	return c.ID
}
