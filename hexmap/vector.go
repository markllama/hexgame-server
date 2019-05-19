// hexmap/vector - a representation of a single location on a hex map
package hexmap

//import (
//	"encoding/json"
//)

type Vector struct {
	Hx int `json:"hx"`
	Hy int `json:"hy"`
}

func (v Vector) Hz() int {
	return v.Hy - v.Hx
}

func (v Vector) Length() int {
	return max(abs(v.Hx), abs(v.Hy), abs(v.Hz()))
}
