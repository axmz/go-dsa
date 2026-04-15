package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	maxLength := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				maxLength = max(maxLength, dp[i][j])
			}
		}
	}

	return maxLength
}

func findLength2(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)

	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i >= m || j >= n {
			return 0
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		if nums1[i] == nums2[j] {
			memo[i][j] = 1 + dfs(i+1, j+1)
		} else {
			memo[i][j] = 0
		}

		return memo[i][j]
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res = max(res, dfs(i, j))
		}
	}

	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
