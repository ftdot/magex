package vector2

// Original code by deeean:
// https://github.com/deeean/go-vector/blob/master/vector2/vector2.go

import (
	"fmt"
	"math"

	"github.com/ftdot/magex/utils/mmath"
)

var (
	Null     = New(0, 0)   // Vector2(0, 0)
	Right    = New(1, 0)   // Vector2(1, 0)
	Left     = New(-1, 0)  // Vector2(-1, 0)
	Up       = New(0, -1)  // Vector2(0, -1)
	Down     = New(0, 1)   // Vector2(0, 1)
	Identity = New(1, 1)   // Vector2(1, 1)
	Negative = New(-1, -1) // Vector2(-1, -1)
)

type Vector2 struct {
	X, Y float64
}

func New(x, y float64) *Vector2 {
	return &Vector2{X: x, Y: y}
}

func NewInt(x, y int) *Vector2 {
	return &Vector2{X: float64(x), Y: float64(y)}
}

func (v *Vector2) Copy() *Vector2 {
	return New(v.X, v.Y)
}

func (v *Vector2) Set(x, y float64) *Vector2 {
	v.X = x
	v.Y = y
	return v
}

func (v *Vector2) Add(other *Vector2) *Vector2 {
	return New(v.X+other.X, v.Y+other.Y)
}

func (v *Vector2) AddScalar(scalar float64) *Vector2 {
	return New(v.X+scalar, v.Y+scalar)
}

func (v *Vector2) AddScalars(x, y float64) *Vector2 {
	return New(v.X+x, v.Y+y)
}

func (v *Vector2) Sub(other *Vector2) *Vector2 {
	return New(v.X-other.X, v.Y-other.Y)
}

func (v *Vector2) SubScalar(scalar float64) *Vector2 {
	return New(v.X-scalar, v.Y-scalar)
}

func (v *Vector2) SubScalars(x, y float64) *Vector2 {
	return New(v.X-x, v.Y-y)
}

func (v *Vector2) Mul(other *Vector2) *Vector2 {
	return New(v.X*other.X, v.Y*other.Y)
}

func (v *Vector2) MulScalar(scalar float64) *Vector2 {
	return New(v.X*scalar, v.Y*scalar)
}

func (v *Vector2) MulScalars(x, y float64) *Vector2 {
	return New(v.X*x, v.Y*y)
}

func (v *Vector2) Div(other *Vector2) *Vector2 {
	return New(v.X/other.X, v.Y/other.Y)
}

func (v *Vector2) DivScalar(scalar float64) *Vector2 {
	return New(v.X/scalar, v.Y/scalar)
}

// Divide vector coordinates by the given values.
func (v *Vector2) DivScalars(x, y float64) *Vector2 {
	return New(v.X/x, v.Y/y)
}

// Returns a new vector rotated counter-clockwise by the specified number of radians.
func (v *Vector2) Rotate(angle float64) *Vector2 {
	return New(
		v.X*math.Cos(angle)-v.Y*math.Sin(angle),
		v.X*math.Sin(angle)+v.Y*math.Cos(angle))
}

// Returns the squared length of this vector
func (v *Vector2) Len2() float64 {
	return v.Dot(v)
}

// Returns the length of this vector
func (v *Vector2) Len() float64 {
	return math.Sqrt(v.Len2())
}

// Returns a new vector perpendicular from this one.
func (v *Vector2) Perp() *Vector2 {
	return New(v.Y, -v.X)
}

// Returns a new vector scaled in the direction of X and Y by value x
func (v *Vector2) Scale(x float64) *Vector2 {
	return New(v.X*x, v.Y*x)
}

// Returns a new vector that is reversed from this one.
func (v *Vector2) Reverse() *Vector2 {
	return Negative.Mul(v)
}

// Returns the distance between a and b.
func (v *Vector2) Distance(other *Vector2) float64 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// Returns dot Product of two vectors.
func (v *Vector2) Dot(other *Vector2) float64 {
	return v.X*other.X + v.Y*other.Y
}

// Linearly interpolates between two points.
func (v *Vector2) Lerp(other *Vector2, t float64) *Vector2 {
	return New(
		v.X+(other.X-v.X)*t,
		v.Y+(other.Y-v.Y)*t,
	)
}

// Returns vector with a magnitude of 1 of this vector.
func (v *Vector2) Normalize() *Vector2 {
	m := v.Len()

	if m > mmath.Epsilon {
		return v.DivScalar(m)
	} else {
		return v.Copy()
	}
}

// Returns a reflected vector off the plane defined by a normal.
func (v *Vector2) Reflect(other *Vector2) *Vector2 {
	factor := -2.0 * v.Dot(other)
	return New(
		factor*v.X+other.X,
		factor*v.Y+other.Y,
	)
}

// Gradually changes a vector towards a desired goal over time.
// The vector is smoothed by some spring-damper like function,
// which will never overshoot. The most common use is for smoothing
// a follow camera.
// Original code by Unity.
func (v *Vector2) SmoothDamp(target, currentVelocity *Vector2, smoothTime, deltaTime float64) *Vector2 {
	return v.SmoothDampMaxSpeed(target, currentVelocity, smoothTime, deltaTime, math.Inf(1))
}

// Gradually changes a vector towards a desired goal over time.
// The vector is smoothed by some spring-damper like function,
// which will never overshoot. The most common use is for smoothing
// a follow camera.
// Original code by Unity.
func (v *Vector2) SmoothDampMaxSpeed(target, currentVelocity *Vector2, smoothTime, deltaTime, maxSpeed float64) *Vector2 {
	smoothTime = max(0.0001, smoothTime)
	omega := 2 / smoothTime

	x := omega * deltaTime
	exp := 1 / (1 + x + 0.48*x*x + 0.235*x*x*x)

	change_x := v.X - target.X
	change_y := v.Y - target.Y

	maxChange := maxSpeed * smoothTime
	maxChangeSq := maxChange * maxChange

	if sqDist := change_x*change_x + change_y*change_y; sqDist > maxChangeSq {
		mag := math.Sqrt(sqDist)
		change_x = change_x / mag * maxChange
		change_y = change_y / mag * maxChange
	}

	t_x := v.X - change_x
	t_y := v.Y - change_y

	temp_x := (currentVelocity.X + omega*change_x) * deltaTime
	temp_y := (currentVelocity.Y + omega*change_y) * deltaTime

	currentVelocity.X = (currentVelocity.X - omega*temp_x) * exp
	currentVelocity.Y = (currentVelocity.Y - omega*temp_y) * exp

	o_x := t_x + (change_x+temp_x)*exp
	o_y := t_y + (change_y+temp_y)*exp

	origMinusCurrent_x := target.X - v.X
	origMinusCurrent_y := target.Y - v.Y
	outMinusOrig_x := o_x - target.X
	outMinusOrig_y := o_y - target.Y

	if origMinusCurrent_x*outMinusOrig_x+origMinusCurrent_y*outMinusOrig_y > 0 {
		o_x = target.X
		o_y = target.Y

		currentVelocity.X = (o_x - target.X) / deltaTime
		currentVelocity.X = (o_y - target.Y) / deltaTime
	}
	return New(o_x, o_y)
}

// Checks whether the vectors are equal.
func (v *Vector2) Equals(other *Vector2) bool {
	return v.X == other.X && v.Y == other.Y
}

// Checks whether this vector is greater than other.
func (v *Vector2) Greater(other *Vector2) bool {
	return v.X > other.X && v.Y > other.Y
}

// Checks whether this vector is greater than or equal
// to other.
func (v *Vector2) GreaterEq(other *Vector2) bool {
	return v.X >= other.X && v.Y >= other.Y
}

// Checks whether this vector is lesser than other.
func (v *Vector2) Lesser(other *Vector2) bool {
	return v.X < other.X && v.Y < other.Y
}

// Checks whether this vector is lesser than or equal
// to other.
func (v *Vector2) LesserEq(other *Vector2) bool {
	return v.X <= other.X && v.Y <= other.Y
}

// Reports whether any vector coordinates is an
// IEEE 754 “not-a-number” value.
func (v *Vector2) IsNaN() bool {
	return math.IsNaN(v.X) || math.IsNaN(v.Y)
}

// IsInf reports whether any vector coordinate is an infinity,
// according to sign.
// If sign > 0, IsInf reports whether any coordinate is positive infinity.
// If sign < 0, IsInf reports whether any coordinate is negative infinity.
// If sign == 0, IsInf reports whether any coordinate is either infinity.
func (v *Vector2) IsInf(sign int) bool {
	return math.IsInf(v.X, sign) || math.IsInf(v.Y, sign)
}

func (v *Vector2) String() string {
	return fmt.Sprintf("Vector2(%f, %f)", v.X, v.Y)
}
