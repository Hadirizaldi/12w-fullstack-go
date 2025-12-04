package leetcode

import (
	"testing"

	"github.com/Hadirizaldi/12w-fullstack-go/leetcode/helpers"
)

func TestIntersectionTwo(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2,2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
		{[]int{1, 2, 3, 3, 3}, []int{3, 4, 5, 3}, []int{3,3}},
		{[]int{}, []int{1, 2, 3}, []int{}},
		{[]int{1, 2, 3}, []int{}, []int{}},
	}
	
	for _, tt := range tests {
		got := IntersectionTwo(tt.nums1, tt.nums2)
		if !helpers.CompareSlicesIgnoringOrder(got, tt.want) {
			t.Errorf("Intersection(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
		}
	}
}