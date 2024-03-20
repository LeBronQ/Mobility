package Mobility

import (
	"testing"
)

func TestBoundCalculation(t *testing.T) {
	box := NewBox(100, 500, 0, 400, 0, 400)
	cases := []struct {
		Name     string
		spx      float64
		initX    float64
		dir      string
		expected float64
	}{
		{"case1", 100, 300, "x", 400},
		{"case2", 300, 300, "x", 400},
		{"case3", 700, 300, "x", 200},
		{"case4", 1600, 300, "x", 300},
		{"case5", 5300, 300, "x", 200},
		{"case6", -100, 300, "x", 200},
		{"case7", -400, 300, "x", 300},
		{"case8", -800, 300, "x", 300},
		{"case9", -3700, 300, "x", 400},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := box.BoundCalculation(c.spx, c.initX, c.dir); ans != c.expected {
				t.Fatalf("expected %f, but %f got", c.expected, ans)
			}
		})
	}
}
