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

var roman map[string]int = map[string]int{
	"III":     3,
	"XIV":     14,
	"MCMXCIV": 1994,
	"LVIII":   58,
	"IX":      9,
	"IV":      4,
	"XCIX":    99,
}

func TestRomanToInt(t *testing.T) {
	for k, v := range roman {
		got := romanToInt(k)
		want := v
		if got != want {
			t.Errorf("romanToInt(%v) wanted: %v, got: %v", k, want, got)
		}
	}
}

var strs0 []string = []string{"flower", "flow", "flight", "flood"}
var strs1 []string = []string{"flower", "dog", "tree"}
var strs2 []string = []string{}
var strs3 []string = []string{"aa", "a"}
var strs4 []string = []string{"dog", "racecar", "car"}

func BenchmarkLcp(b *testing.B) {
	b.StartTimer()

	for i := 1; i < b.N; i++ {
		longestCommonPrefix(strs0)
		// longestCommonPrefix(strs2)
		// longestCommonPrefix(strs3)
		// longestCommonPrefix(strs4)
	}
}
func BenchmarkLcpDC(b *testing.B) {
	b.StartTimer()

	for i := 1; i < b.N; i++ {
		lcpDC(strs0)
		// lcpDC(strs2)
		// lcpDC(strs3)
		// lcpDC(strs4)
	}
}
func BenchmarkLcpBS(b *testing.B) {
	b.StartTimer()

	for i := 1; i < b.N; i++ {
		lcpBS(strs0)
		// lcpBS(strs2)
		// lcpBS(strs3)
		// lcpBS(strs4)
	}
}
