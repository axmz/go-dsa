package main

import "fmt"

func countBits(n int) int {
	count := 0
	for n > 0 {
		count++
		n &= (n - 1)
	}
	return count
}

func sumIndicesWithKSetBits(nums []int, k int) int {
	sum := 0
	for i, n := range nums {
		if countBits(i) == k {
			sum += n
		}
	}
	return sum
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
