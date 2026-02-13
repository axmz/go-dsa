package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxSubarraySumCircular(nums []int) int {
	l := len(nums)
	minSum := nums[0]
	curMinSum := nums[0]
	maxSum := nums[0]
	curMaxSum := nums[0]
	total := nums[0]
	for i := 1; i < l; i++ {
		total += nums[i]
		curMaxSum = max(nums[i], curMaxSum+nums[i])
		maxSum = max(maxSum, curMaxSum)
		curMinSum = min(nums[i], curMinSum+nums[i])
		minSum = min(minSum, curMinSum)
	}

	if total == minSum {
		return maxSum
	}
	return max(maxSum, total-minSum)
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
