package main

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		one, _ := strconv.Atoi(s[i-1 : i])
		two, _ := strconv.Atoi(s[i-2 : i])

		if one >= 1 {
			dp[i] += dp[i-1]
		}
		if two >= 10 && two <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
