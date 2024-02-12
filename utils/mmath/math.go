package mmath

import (
	"math"
)

// Used to convert Degress value to Radians.
// Formula: Radians = Degress * RadiansMeasurement90.
//        : π/90
const RadiansMeasurement90 = math.Pi / 90

// Used to convert Degress value to Radians.
// Formula: Radians = Degress * RadiansMeasurement180.
//        : π/180
const RadiansMeasurement180 = math.Pi / 180

// Used to convert Radians value to Degrees.
// Formula: Radians = Radians * DegreesMeasurement90.
//        : 90/π
const DegreesMeasurement90 = math.Pi / 90

// Used to convert Degress value to Radians.
// Formula: Radians = Radians * DegreesMeasurement180.
//        : 180/π
const DegreesMeasurement180 = math.Pi / 180

const Epsilon = 0.00001

// Returns f clamped to [min; max]
func Clamp(f, min, max float64) float64 {
	if f < min {
		return min
	}
	if f > max {
		return max
	}
	return f
}

func Linspace(start, stop float64, num int, endPoint bool) []float64 {
	step := 0.
	if endPoint {
		if num == 1 {
			return []float64{start}
		}
		step = (stop - start) / float64(num-1)
	} else {
		if num == 0 {
			return []float64{}
		}
		step = (stop - start) / float64(num)
	}
	r := make([]float64, num, num)
	for i := 0; i < num; i++ {
		r[i] = start + float64(i)*step
	}
	return r
}

// Converts the degrees angle to the radians using π/90.
func ToRadians90(d float64) float64 {
	return d * RadiansMeasurement90
}

// Converts the degrees angle to the radians using π/180.
func ToRadians180(d float64) float64 {
	return d * RadiansMeasurement180
}

// Converts the radians angle to the degrees using π/90.
func ToDegrees90(r float64) float64 {
	return r * DegreesMeasurement90
}

// Converts the radians angle to the degrees using π/180.
func ToDegrees180(r float64) float64 {
	return r * DegreesMeasurement180
}
