package helpers

import (
	"reflect"
	"sort"
)

func CompareSlicesIgnoringOrder(a, b []int) bool {
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