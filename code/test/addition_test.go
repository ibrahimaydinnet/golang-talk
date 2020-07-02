package main

import "testing"

func TestAddition(t *testing.T) {
	cases := []struct {
		x, y, want int
	}{
		{1, 2, 3},
		{2, 3, 5},
	}

	for _, c := range cases {
		ans := c.x + c.y
		if ans != c.want {
			t.Errorf("want %d, got %d", c.want, ans)
		}
	}
}
