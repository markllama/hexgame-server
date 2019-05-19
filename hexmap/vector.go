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

func (v0 Vector) Add(v1 Vector) Vector {
	return Vector{v0.Hx + v1.Hx, v0.Hy + v1.Hy}
}
