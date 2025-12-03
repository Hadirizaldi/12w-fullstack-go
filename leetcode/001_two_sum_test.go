package leetcode

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	} {
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 2, 3, 4, 5}, 9, []int{3, 4}},
		{[]int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
	}

	for _, tt := range tests {
		got := TwoSum(tt.nums, tt.target)
		if !sliceEqual(got, tt.want) {
			t.Errorf("TwoSum(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for index := range a {
		if a[index] != b[index] {
			return false
		}
	}

	return true
}