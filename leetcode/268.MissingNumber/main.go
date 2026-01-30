package main

import "fmt"

func missingNumber(nums []int) int {
	xor := 0
	for i := 0; i < len(nums); i++ {
		xor ^= i ^ nums[i]
	}
	xor ^= len(nums) // there is one less number in nums than in range [0, n]
	return xor
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
