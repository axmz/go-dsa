package main

import "fmt"

func canJump2(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[len(nums)-1] = true
	for i := len(nums) - 2; i >= 0; i-- {
		for j := i + 1; j <= i+nums[i] && j < len(nums); j++ {
			if dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func canJump(nums []int) bool {
	reach := 0
	for i := 0; i < len(nums); i++ {
		if i > reach {
			return false
		}
		reach = max(reach, i+nums[i])
	}

	return true
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
