package main

import "testing"

func TestMultiplication(t *testing.T) {
	cases := []struct {
		x, y, want int
	}{
		{1, 2, 2},
		{2, 3, 6},
	}

	for _, c := range cases {
		ans := c.x * c.y
		if ans != c.want {
			t.Errorf("want %d, got %d", c.want, ans)
		}
	}
}
