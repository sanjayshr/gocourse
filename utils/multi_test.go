package utils

import "testing"

func TestSwap(t *testing.T) {
	x, y := Swap(2, 4)
	a, b := 4, 2

	if x != a && y != b {
		t.Errorf("x:%d y:%d, a:%d b:%d", x, y, a, b)
	}
}
