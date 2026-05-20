package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func stoneGameII(piles []int) int {
	memo := map[[2]int]int{}

	prefixSum := make([]int, len(piles)+1)
	for i := 0; i < len(piles); i++ {
		prefixSum[i+1] = prefixSum[i] + piles[i]
	}

	var dfs func(i, m int) int
	dfs = func(i, m int) int {
		if i >= len(piles) {
			return 0
		}

		key := [2]int{i, m}
		if v, ok := memo[key]; ok {
			return v
		}

		res := -1 << 30
		for x := 1; x <= 2*m; x++ {
			if i+x-1 >= len(piles) {
				break
			}
			sum := prefixSum[i+x] - prefixSum[i]
			res = max(res, sum-dfs(i+x, max(m, x)))
		}
		memo[key] = res
		return res
	}

	diff := dfs(0, 1)
	total := prefixSum[len(piles)]
	return (total + diff) / 2
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
