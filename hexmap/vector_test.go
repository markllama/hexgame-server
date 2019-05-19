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

	hv1 := Vector{1, 1}
	if hv1.Length() != 1 {
		t.Error()
	}
	
	hv2 := Vector{-1, -1}
	if hv2.Length() != 1 {
		t.Error()
	}
	
	hv3 := Vector{0, -1}
	if hv3.Length() != 1 {
		t.Error()
	}
}

func TestAdd(t *testing.T) {
	hv0 := Vector{5, 6}
	hv1 := Vector{5, 5}

	result0 := hv0.Add(ORIGIN)

	if result0 != hv0 {
		t.Error()
	}

	result1 := hv0.Add(UNIT[0])

	if result1 != hv1 {
		t.Errorf("hv0 %s != result1 %s", hv1, result1)
	}
}
