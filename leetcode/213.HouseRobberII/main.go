package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func robLinear(nums []int) int {
	prevTwo, prevOne := 0, 0
	for _, value := range nums {
		current := max(prevOne, prevTwo+value)
		prevTwo = prevOne
		prevOne = current
	}
	return prevOne
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(robLinear(nums[:len(nums)-1]), robLinear(nums[1:]))
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
