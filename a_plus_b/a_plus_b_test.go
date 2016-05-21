package a_plus_b

import (
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		a, b int
	}{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	want := []int{3, 7, 11}
	for i, c := range cases {
		got := Add(c.a, c.b)
		if got != want[i] {
			t.Errorf("Add(%v, %v) == %v, want %v", c.a, c.b, got, want[i])
		}
	}
}
