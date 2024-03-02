package collision2d

import (
	"fmt"
	"math"

	"github.com/ftdot/magex/utils/mmath"
	"github.com/ftdot/magex/utils/vector2"
)

const defaultResolution = 100

////

// Circle is a struct that represents a circle with a position and a raidus.
type Circle struct {
	Pos        *vector2.Vector2
	R          float64
	SizeScalar *vector2.Vector2
	Resolution int
}

func (circle Circle) String() string {
	return fmt.Sprintf("{Pos:%s Radius: %f SizeScalar: %s}", circle.Pos, circle.R, circle.SizeScalar)
}

// NewCircle create a new circle with vector pos as center and radius r
func NewCircle(pos *vector2.Vector2, r float64, sizeScalar *vector2.Vector2) *Circle {
	return &Circle{
		Pos: pos,
		R: r,
		Resolution: defaultResolution,
		SizeScalar: sizeScalar,
	}
}

// GetAABB returns the axis-aligned bounding box of the circle.
func (circle *Circle) GetAABB() *Polygon {
	r := circle.R
	corner := circle.Pos.Sub(vector2.New(r, r))
	polygon := NewBox(corner, r*2, r*2).ToPolygon()
	return polygon
}

func (circle *Circle) ToPolygon() *Polygon {
	pointsN := 2*circle.Resolution - 1
	points := make([]float64, pointsN+1)
	thetaL := mmath.Linspace(0, 2*math.Pi, circle.Resolution, false)
	
	for i := 0; i < pointsN; i += 2 {
		t := thetaL[i/2]
		points[i] = math.Cos(t) * circle.R * circle.SizeScalar.X
		points[i+1] = math.Sin(t) * circle.R * circle.SizeScalar.Y
	}

	return NewPolygon(circle.Pos, vector2.Null.Copy(), 0, points)
}
