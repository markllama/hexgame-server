package hexmap

import "testing"

func TestVector(t *testing.T) {
	hv := Vector{Hx: 0, Hy: 0}
	
	if hv.Hx != 0 {
		t.Error()
	}
	
	if hv.Hy != 0 {
		t.Error()
	}
}

func TestLength(t *testing.T) {
	hv0 := Vector{Hx: 0, Hy: 0}

	if hv0.Length() != 0 {
		t.Error()
	}
}
