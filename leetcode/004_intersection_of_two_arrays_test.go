package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// Catatan: Fungsi Intersection (yang Anda buat) dan mapTolist harus berada di package main

// Fungsi pembanding yang benar untuk slice yang tidak terurut (Order-Agnostic Comparison)
func compareSlicesIgnoringOrder(a, b []int) bool {
    // 1. Cek Panjang
    if len(a) != len(b) {
        return false
    }
    
    // 2. Urutkan kedua slice
    // Ini adalah cara standar untuk membandingkan slice yang tidak peduli urutan
    sort.Ints(a)
    sort.Ints(b)

    // 3. Bandingkan kontennya yang sudah diurutkan (DeepEqual aman untuk slice terurut)
    return reflect.DeepEqual(a, b)
}


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
        if !compareSlicesIgnoringOrder(got, tt.want) {
            // Kita harus membuat salinan slice untuk di-sort saat error message
            // karena sorting di dalam compareSlicesIgnoringOrder mengubah slice aslinya.
            t.Errorf("Intersection(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
        }
    }
}