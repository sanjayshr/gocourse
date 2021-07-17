package utils

import (
	"testing"
)

func TestDeepEQL(t *testing.T) {

	t.Run("Test String", func(t *testing.T) {
		m1 := map[int]string{
			1: "p1",
			2: "p2",
		}

		m2 := map[int]string{
			1: "p1",
			2: "p2",
		}

		got := DeepEQLString(m1, m2)
		want := true

		if got != want {
			t.Errorf("got: %v want: %v", got, want)
		}

	})

	t.Run("Test Integers", func(t *testing.T) {
		m1 := map[int]int{
			1: 1,
			2: 2,
		}

		m2 := map[int]int{
			1: 1,
			2: 2,
		}

		got := DeepEQLInt(m1, m2)
		want := true

		if got != want {
			t.Errorf("got: %v want: %v", got, want)
		}

	})

}
