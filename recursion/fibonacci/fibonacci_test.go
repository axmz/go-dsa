package fibonacci_test

import (
	fibonacci "godsa/recursion/fibonacci"
	"testing"
)

var data int = 13

func BenchmarkIfibonacci1(b *testing.B) {
	// b.StartTimer()
	for i := 0; i < b.N; i++ {
		fibonacci.Ifibonacci1(data)
	}
}
func BenchmarkIfibonacci2(b *testing.B) {
	// b.StartTimer()
	for i := 0; i < b.N; i++ {
		fibonacci.Ifibonacci1(data)
	}
}
func BenchmarkIfibonacci3(b *testing.B) {
	// b.StartTimer()
	for i := 0; i < b.N; i++ {
		fibonacci.Ifibonacci1(data)
	}
}

func BenchmarkRfibonacci(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		fibonacci.Rfibonacci(data)
	}
}
