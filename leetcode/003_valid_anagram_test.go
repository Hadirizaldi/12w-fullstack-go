package leetcode

import "testing"

func TestValidAnagram(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"", "", true},
		{"a", "ab", false},
		{"listen", "silent", true},
		{"hello", "billion", false},
	}

	for _, tt := range tests {
		got := IsAnagram(tt.s, tt.t)
		if got != tt.want {
			t.Errorf("IsAnagram(%q, %q) = %v; want %v", tt.s, tt.t, got, tt.want)
		}
	}
}