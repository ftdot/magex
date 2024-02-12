package rectutil

import "github.com/ftdot/magex/utils/vector2"

// Returns true if point is on rect
func IsPointOnRect(
	point *vector2.Vector2,
	rectStart *vector2.Vector2,
	rectEnd *vector2.Vector2) bool {
	return point.Greater(rectStart) && point.Lesser(rectEnd)
}

// Returns true if rect is collides with rect
func IsRectCollidesRect(
	rect1Start *vector2.Vector2,
	rect1End *vector2.Vector2,
	rect2Start *vector2.Vector2,
	rect2End *vector2.Vector2,
) bool {
	return rect1Start.GreaterEq(rect2Start) && rect1End.Lesser(rect2End)
}
