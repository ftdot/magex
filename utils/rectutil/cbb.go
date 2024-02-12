package rectutil

import (
	"math"

	"github.com/ftdot/magex/utils/mmath"
	"github.com/ftdot/magex/utils/vector2"
)

// Computes the bounding box of rotated rectangle.
// x, y - coordinates of the rectangle
// w, h - width and height of the rectangle
// d - angle in degress
func ComputeBoundingBox(x, y, w, h, d float64) (A, B *vector2.Vector2) {
	a := math.Mod(d, 90) * mmath.RadiansMeasurement90
	ux := math.Cos(a) * 0.5
	uy := math.Sin(a) * 0.5
	wx := w * ux
	wy := w * uy
	hx := h * -uy
	hy := h * ux

	x1 := x - wx - hx
	y1 := y - wy - hy
	x2 := x + wx - hx
	y2 := y + wy - hy
	x3 := x + wx + hx
	y3 := y + wy + hy
	x4 := x - wx + hx
	y4 := y - wy + hy

	return vector2.New(
			min(x1, x2, x3, x4),
			min(y1, y2, y3, y4),
		), vector2.New(
			max(x1, x2, x3, x4),
			max(y1, y2, y3, y4),
		)
}
