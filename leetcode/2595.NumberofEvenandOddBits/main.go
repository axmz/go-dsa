package main

import "fmt"

func evenOddBit(n int) []int {
	even, odd := 0, 0
	for n > 0 {
		even += n & 1
		n >>= 1
		// odd position, not byte responsible for the number being odd
		odd += n & 1
		n >>= 1
	}

	return []int{even, odd}
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
