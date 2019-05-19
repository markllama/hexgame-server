package hexmap

import "testing"

func TestORIGIN(t *testing.T) {
	if ORIGIN.Hx != 0 {
		t.Error()
	}
	if ORIGIN.Hy != 0 {
		t.Error()
	}
}

//func TestUnit(t *testing.T) {
//	
//}
