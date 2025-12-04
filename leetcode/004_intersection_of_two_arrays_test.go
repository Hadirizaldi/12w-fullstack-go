package leetcode

import (
	"testing"

	"github.com/Hadirizaldi/12w-fullstack-go/leetcode/helpers"
)


func TestIntersection(t *testing.T) {
    tests := []struct {
        nums1 []int
        nums2 []int
        want  []int
    }{
        {[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
        {[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}}, 
        {[]int{1, 2, 3}, []int{4, 5, 6}, []int{}},
        {[]int{}, []int{1, 2, 3}, []int{}},
        {[]int{1, 2, 3}, []int{}, []int{}},
    }

    for _, tt := range tests {
        got := Intersection(tt.nums1, tt.nums2)

        // GUNAKAN FUNGSI PEMBANDING KHUSUS
        if !helpers.CompareSlicesIgnoringOrder(got, tt.want) {
            // Kita harus membuat salinan slice untuk di-sort saat error message
            // karena sorting di dalam compareSlicesIgnoringOrder mengubah slice aslinya.
            t.Errorf("Intersection(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
        }
    }
}