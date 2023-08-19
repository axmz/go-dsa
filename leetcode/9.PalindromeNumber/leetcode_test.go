package main

import (
	"testing"
)

var palindrome map[int]bool = map[int]bool{
	1221: true,
	121:  true,
	-121: false,
	120:  false,
	0:    false,
}

func TestIsPalindrome(t *testing.T) {
	for k, v := range palindrome {
		got := isPalindrome(k)
		want := v
		if got != want {
			t.Errorf("isPalindrome(%v) wanted: %v, got: %v", k, want, got)
		}
	}
}
