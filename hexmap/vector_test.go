package hexmap

import (
	"testing"
	"fmt"
	"encoding/json"
)

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

func TestDistance(t *testing.T) {
	hv0 := Vector{-3, 5}
	hv1 := Vector{3, -5}

	if ORIGIN.Distance(hv0) != hv0.Length() {
		t.Error()
	}

	if ORIGIN.Distance(hv1) != hv1.Length() {
		t.Error()
	}
}

func TestRotate(t *testing.T) {

	// check that rotating the unit vectors gives the right answers
	for i, v := range UNIT[:5] {
		if UNIT[0].Rotate(i) != v {
			t.Errorf("%s rotated 1 is %s, not %s", UNIT[0], UNIT[0].Rotate(i), UNIT[i])
		}
	}

	// try off-axis
	ring0 := []Vector {
		{3, -4},
		{7, 3},
		{4, 7},
		{-3, 4},
		{-7, -3},
		{-4, -7},
	}

	for i, v := range ring0 {
		if ring0[0].Rotate(i) != v {
			t.Errorf("%s rotated %d is %s, not %s", ring0[0], i, ring0[0].Rotate(i), ring0[i])
		}
	}
	
}

func TestJson(t *testing.T) {

	originJson := `{"hx":0,"hy":0}`

	// try marshaling
	o, err := json.Marshal(ORIGIN)

	if err != nil {
		t.Errorf(fmt.Sprintf("%s", err))
	}

	if string(o) != originJson {
		t.Errorf("Marshal ORIGIN didnt work: %s != %s", originJson, string(o))
	}

	// now unmarshal

	var originVector Vector
	err = json.Unmarshal([]byte(originJson), &originVector)
	
	if err != nil {
		t.Errorf(fmt.Sprintf("%s", err))
	}

	if originVector != ORIGIN {
		t.Errorf("Origin Vector JSON did not unmarshal to ORIGIN")
	}
	
}
