package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProduct(nums []int) int {
	maxMax := nums[0]
	curMax, curMin := nums[0], nums[0]

	for _, n := range nums[1:] {
		prevMax, prevMin := curMax, curMin
		curMax = max(n, max(prevMax*n, prevMin*n))
		curMin = min(n, min(prevMax*n, prevMin*n))
		maxMax = max(maxMax, curMax)
	}

	return maxMax
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
