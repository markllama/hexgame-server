package hexmap

var ORIGIN Vector
var UNIT [7]Vector

func init() {
	// The location of the metric
	ORIGIN = Vector{}

	// 6 unit vectors and one extra to make some algorithms work better.
	UNIT = [7]Vector {
		Vector{0,  -1},
		Vector{1,   0},
		Vector{1,   1},
		Vector{0,   1},
		Vector{-1,  0},
		Vector{-1, -1},
		Vector{0,  -1},
	}
}
