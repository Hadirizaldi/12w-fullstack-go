package leetcode

import "testing"

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0, 0, 1}, []int{1, 0, 0}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{}, []int{}},
	}

	for _, tt := range tests {
		got := MoveZeroes(tt.nums)
		for i := range got {
			if got[i] != tt.want[i] {
				t.Errorf("MoveZeroes(%v) = %v; want %v", tt.nums, got, tt.want)
				break
			}
		}
	}
}