package leetcode

import "testing"

func TestContainsDuplicate(t *testing.T) {

	tests := []struct {
		nums []int
		want bool
	} {
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
		{[]int{}, false},
		{[]int{1}, false},
		{[]int{1, 2, 2, 3, 4, 5, 6}, true},
	}

	for _, tt := range tests {
		got := ContainsDulicate(tt.nums)
		if got != tt.want {
			t.Errorf("ContainsDulicate(%v) = %v; want %v", tt.nums, got, tt.want)
		}
	}
}