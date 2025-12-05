package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{5, 4, 3, 2, 1}, 0},
		{[]int{3, 2, 6, 5, 0, 3}, 4},
		{[]int{}, 0},
	}

	for _, tt := range tests {
		got := MaxProfit(tt.prices)
		if got != tt.want {
			t.Errorf("MaxProfit(%v) = %d; want %d", tt.prices, got, tt.want)
		}
	}
}