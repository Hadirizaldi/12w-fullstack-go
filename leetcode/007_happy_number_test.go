package leetcode

import "testing"

func TestIsHappy(t *testing.T) {

	tests := []struct {
		n    int
		want bool
	}{
		{19, true},
		{2, false},
		{1, true},
		{7, true},
		{20, false},
	}
	
	for _, tt := range tests {
		got := IsHappy(tt.n)
		if got != tt.want {
			t.Errorf("IsHappy(%v) = %v; want %v", tt.n, got, tt.want)
		}
	}
}