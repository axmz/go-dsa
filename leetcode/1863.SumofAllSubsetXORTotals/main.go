package main

import "fmt"

func subsetXORSum(nums []int) int {
	xor := 0
	for _, num := range nums {
		xor |= num
	}
	return xor * (1 << (len(nums) - 1))
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
