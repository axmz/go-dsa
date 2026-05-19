package main

import "fmt"

func combinationSum4_2(nums []int, target int) int {
	var dfs func(int) int
	memo := make(map[int]int)

	dfs = func(remaining int) int {
		if remaining == 0 {
			return 1
		}
		if val, exists := memo[remaining]; exists {
			return val
		}

		count := 0
		for _, num := range nums {
			if num <= remaining {
				count += dfs(remaining - num)
			}
		}

		memo[remaining] = count
		return count
	}

	return dfs(target)
}

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}

	return dp[target]
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
