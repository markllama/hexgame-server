// hexmap/vector - a representation of a single location on a hex map
package hexmap

import (
	"fmt"
//	"encoding/json"
)

type Vector struct {
	Hx int `json:"hx"`
	Hy int `json:"hy"`
}

func (v Vector) String() string {
	return fmt.Sprintf("Vector(hx: %d, hy: %d)", v.Hx, v.Hy)
}

// Hz is dependent on Hx and Hy
func (v Vector) Hz() int {
	return v.Hy - v.Hx
}

// The length of a HexVector is the absolute value of the largest component
func (v Vector) Length() int {
	return max(abs(v.Hx), abs(v.Hy), abs(v.Hz()))
}

// Reverse the direction of a vector.
func (v Vector) Invert() Vector {
	return Vector{-v.Hx, -v.Hy}
}

// Adding two vectors is just adding the components
func (v0 Vector) Add(v1 Vector) Vector {
	return Vector{v0.Hx + v1.Hx, v0.Hy + v1.Hy}
}

// Subtracting two vectors is the same as adding one to the inverse of the other
func (v0 Vector) Sub(v1 Vector) Vector {
	return v0.Add(v1.Invert())
}

// The distance between two points is the length of the difference of the
// vectors
func (v0 Vector) Distance(v1 Vector) int {
	return v1.Sub(v0).Length()
}

// Hex maps form a mathematical "ring".  Each hex has 6 "cardinal points"
// around the origin, in the same way cartisian coordinates have 4.
func (v Vector) Rotate(hextants int) Vector {
	// the hextants can be any int.  first reduce it to range [0..5]

	hextants = hextants % 6

	if hextants < 0 {
		hextants += 6
	}
	
	ring := []int {v.Hx, v.Hy, v.Hz(), -v.Hx, -v.Hy, -v.Hz(), v.Hx}

	return Vector{ring[hextants], ring[hextants+1]}
}
