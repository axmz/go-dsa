package main

import (
	"fmt"
	"math/rand"
	"testing"

	"godsa/sorting/bubble"
	"godsa/sorting/bucket"
	"godsa/sorting/counting"
	"godsa/sorting/heap"
	"godsa/sorting/insertion"
	"godsa/sorting/merge"
	"godsa/sorting/quick"
	"godsa/sorting/radix"
	"godsa/sorting/selection"
)

func genBaseWithSeed(seed int64, n int) []int {
	rng := rand.New(rand.NewSource(seed))
	s := make([]int, n)
	for i := range s {
		s[i] = rng.Intn(n * 10)
	}
	return s
}

var sorts = map[string]func([]int) []int{
	"bubble":    bubble.Sort,
	"selection": selection.Sort,
	"insertion": insertion.Sort,
	"radix":     radix.Sort,
	"merge":     merge.Sort,
	"quick":     quick.Sort,
	"heap":      heap.Sort,
	"counting":  counting.Sort,
	"bucket":    bucket.Sort,
}

func BenchmarkSorts(b *testing.B) {
	// Use a single size for all algorithms to keep comparisons simple
	n := 800
	b.ReportAllocs()
	base := genBaseWithSeed(42, n)

	for name, fn := range sorts {
		b.Run(fmt.Sprintf("N=%d/%s", n, name), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				a := make([]int, len(base))
				copy(a, base)
				fn(a)
			}
		})
	}
}
