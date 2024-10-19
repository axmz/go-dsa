package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func robBottomUp(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[len(dp)-1]
}

func robTopDown(nums []int) int {
	memo := make([]int, len(nums))
	for i := range memo {
		memo[i] = -1
	}

	var dp func(int) int
	dp = func(i int) int {
		res := 0

		if memo[i] != -1 {
			return memo[i]
		}

		if i == 0 {
			res = nums[0]
		} else if i == 1 {
			res = max(nums[0], nums[1])
		} else {
			res = max(nums[i]+dp(i-2), dp(i-1))
		}

		memo[i] = res
		return res
	}

	return dp(len(nums) - 1)
}

func main() {
	// houses := []int{1, 2, 3, 1}
	// houses := []int{2, 7, 9, 3, 1}
	// houses := []int{2, 1}
	houses := []int{2, 1, 1, 2}

	fmt.Println(robTopDown(houses))
	fmt.Println(robBottomUp(houses))
}
