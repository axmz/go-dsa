package main

import "fmt"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lastStoneWeightII(stones []int) int {
	memo := map[[2]int]int{}

	var dfs func(i, diff int) int
	dfs = func(i, diff int) int {
		if i == len(stones) {
			return abs(diff)
		}

		key := [2]int{i, diff}
		if v, ok := memo[key]; ok {
			return v
		}

		add := dfs(i+1, diff+stones[i])
		sub := dfs(i+1, diff-stones[i])
		memo[key] = min(add, sub)
		return memo[key]
	}

	return dfs(0, 0)
}

func lastStoneWeightII2(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}

	dp := make([]bool, sum+1)
	dp[0] = true
	for _, v := range stones {
		for j := sum; j >= v; j-- {
			dp[j] = dp[j] || dp[j-v]
		}
	}

	for i := sum / 2; i >= 0; i-- {
		if dp[i] {
			return sum - 2*i
		}
	}
	return 0
}

func lastStoneWeightII3(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}
	target := sum / 2

	memo := map[[2]int]int{}

	var dfs func(i, curr int) int
	dfs = func(i, curr int) int {
		if curr > target {
			return -1
		}
		if i == len(stones) {
			return curr
		}

		key := [2]int{i, curr}
		if v, ok := memo[key]; ok {
			return v
		}

		take := dfs(i+1, curr+stones[i])
		skip := dfs(i+1, curr)
		if take > skip {
			memo[key] = take
		} else {
			memo[key] = skip
		}
		return memo[key]
	}

	best := dfs(0, 0)
	return sum - 2*best
}

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeightII(stones))
}
