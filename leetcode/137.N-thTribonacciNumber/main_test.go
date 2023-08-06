package main

import (
	"testing"
)

var num = 25

func BenchmarkTribonacci3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tribonacci3(num)
	}
}

func BenchmarkTribonacci2(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		tribonacci2(num)
	}
}

func BenchmarkTribonacci(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tribonacci(num)
	}
}
