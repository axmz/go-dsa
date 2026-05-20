package main

import "fmt"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func stoneGame(piles []int) bool {
	memo := map[[2]int]int{}

	var dfs func(l, r int) int
	dfs = func(l, r int) int {
		if l > r {
			return 0
		}

		key := [2]int{l, r}
		if v, ok := memo[key]; ok {
			return v
		}

		left := abs(dfs(l+1, r) + piles[l])
		right := abs(dfs(l, r-1) + piles[r])

		memo[key] = max(left, right)
		return memo[key]
	}

	if dfs(0, len(piles)-1) > 0 {
		return true
	}
	return false
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
