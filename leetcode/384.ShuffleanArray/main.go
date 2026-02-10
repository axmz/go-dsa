package main

import (
	"fmt"
	"math/rand"
)

// Two methods to shuffle an array
// 1. Using built-in rand.Shuffle
// 2. Using Fisher-Yates Algorithm
type Solution struct {
	shuffled []int
	orig     []int
}

func Constructor(nums []int) Solution {
	perm := make([]int, len(nums))
	copy(perm, nums)
	return Solution{
		shuffled: perm,
		orig:     nums,
	}
}

func (this *Solution) Reset() []int {
	copy(this.shuffled, this.orig)
	return this.shuffled
}

func (this *Solution) Shuffle() []int {
	// Shuffle in-place
	// rand.Shuffle(len(this.perm), func(i, j int) {
	// 	this.perm[i], this.perm[j] = this.perm[j], this.perm[i]
	// })

	// Fisher-Yates Algorithm
	for i := range this.shuffled {
		j := i + rand.Intn(len(this.shuffled)-i)
		this.shuffled[i], this.shuffled[j] = this.shuffled[j], this.shuffled[i]
	}
	// for i := len(this.perm) - 1; i > 0; i-- {
	// 	j := rand.Intn(i + 1)
	// 	this.perm[i], this.perm[j] = this.perm[j], this.perm[i]
	// }
	return this.shuffled
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
