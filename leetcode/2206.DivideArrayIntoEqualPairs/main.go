package main

import "fmt"

// Nice 64+ bitset solution
func divideArray(nums []int) bool {
	// instead of hardcoding max ex: 500 as in constraints
	// we can calculate max from input
	max := 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	buckets := max/64 + 1
	counts := make([]uint64, buckets)
	for _, num := range nums {
		bucket := num / 64
		pos := num % 64
		counts[bucket] ^= 1 << pos // there was no need to check if 0 or 1, xor does the job
	}
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}
	return true
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
