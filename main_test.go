package main

import "testing"

func TestMaxRevenue(t *testing.T) {
	classic := []int{1, 5, 8, 9, 10, 17, 17, 20}
	cases := []struct {
		prices []int
		n      int
		want   int
	}{
		{classic, 8, 22}, // 2 + 6
		{classic, 4, 10}, // 2 + 2
		{classic, 1, 1},
		{classic, 0, 0},
		{[]int{3}, 5, 15},     // only length-1 pieces exist
		{[]int{1, 10}, 5, 21}, // 2+2+1
		{nil, 4, 0},
	}
	for _, c := range cases {
		if got := MaxRevenue(c.prices, c.n); got != c.want {
			t.Errorf("MaxRevenue(%v,%d) = %d, want %d", c.prices, c.n, got, c.want)
		}
	}
}
