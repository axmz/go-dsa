package main

import (
	"testing"
)

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
