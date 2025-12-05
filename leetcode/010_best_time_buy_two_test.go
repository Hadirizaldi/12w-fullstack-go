package leetcode

import "testing"

func TextMaxProfitTwo(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{5, 4, 3, 2, 1}, 0},
		{[]int{3, 2, 6, 5, 0, 3}, 7},
		{[]int{}, 0},
	}

	for _, tt := range tests {
		got := MaxProfitTwo(tt.prices)
		if got != tt.want {
			t.Errorf("MaxProfitTwo(%v) = %d; want %d", tt.prices, got, tt.want)
		}
	}
}