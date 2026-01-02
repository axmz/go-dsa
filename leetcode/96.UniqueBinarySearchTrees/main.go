package main

import "fmt"

func numTrees2(n int) int {
	var C int = 1
	for i := 0; i < n; i++ {
		C = C * 2 * int(2*i+1) / int(i+2)
	}
	return int(C)
}

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}

func main() {
	x := 5
	fmt.Println(numTrees(x))
}
