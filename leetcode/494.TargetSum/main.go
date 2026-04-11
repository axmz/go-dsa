package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	total_sum := 0
	for _, num := range nums {
		total_sum += num
	}

	if total_sum < target || (total_sum-target)%2 != 0 {
		return 0
	}

	subset_sum := (total_sum - target) / 2

	dp := make([]int, subset_sum+1)
	dp[0] = 1

	for _, num := range nums {
		for j := subset_sum; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}

	return dp[subset_sum]
}

func findTargetSumWays2(nums []int, target int) int {
	dp := make(map[int]int, target*2+1)
	dp[0] = 1
	for _, num := range nums {
		next := make(map[int]int, target*2+1)
		for sum, count := range dp {
			next[sum+num] += count
			next[sum-num] += count
		}
		dp = next
	}
	return dp[target]
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
