package collision2d

import (
	"fmt"
	"math"

	"github.com/ftdot/magex/utils/vector2"
)

// Response contains the information about an collision test.
type Response struct {
	A, B               interface{}
	Overlap            float64
	OverlapN, OverlapV *vector2.Vector2
	AInB, BInA         bool
}

func (response Response) String() string {
	output := string("Response:\n{A: %s\nB: %s\nOverlap: %f\nOverlapN: %sOverlapV: %sAInB: %t, BInA: %t}")
	return fmt.Sprintf(output, response.A, response.B, response.Overlap, response.OverlapN, response.OverlapV, response.AInB, response.BInA)
}

// NewResponse is used to create a new response when necessary.
func NewResponse() *Response {
	return &Response{Overlap: math.MaxFloat64, AInB: true, BInA: true}
}

// NotColliding is to be used when A and B are not colliding and response should be ignored.
func (response *Response) NotColliding() *Response {
	return &Response{Overlap: -math.MaxFloat64, AInB: false, BInA: false}
}
