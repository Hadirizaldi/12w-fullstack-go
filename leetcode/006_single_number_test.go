package leetcode

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
		{[]int{0, 1, 0}, 1},
		{[]int{-1, -1, -2}, -2},
	}

	for _, tt := range tests {
		got := SingleNumber(tt.nums)
		if got != tt.want {
			t.Errorf("SingleNumber(%v) = %v; want %v", tt.nums, got, tt.want)
		}
	}
}