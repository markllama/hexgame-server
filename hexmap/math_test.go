package hexmap

import "testing"

func TestMaxN(t *testing.T) {
	if max(1, 2, 3) != 3 {
		t.Error()
	}

	if max(3, 2, 1) != 3 {
		t.Error()
	}
	if max(-2, 0, 2) != 2 {
		t.Error()
	}
}

func TestAbs(t *testing.T) {
	if abs(-2) != 2 {
		t.Error()
	}
	if abs(2) != 2 {
		t.Error()
	}

	if abs(0) != 0 {
		t.Error()
	}
}
