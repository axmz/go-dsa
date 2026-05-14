package main

import "fmt"

func minExtraChar(s string, dictionary []string) int {
	dict := make(map[string]bool)
	for _, word := range dictionary {
		dict[word] = true
	}

	n := len(s)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1 // assume the current character is extra
		for j := 0; j < i; j++ {
			if dict[s[j:i]] {
				if dp[j] < dp[i] {
					dp[i] = dp[j]
				}
			}
		}
	}

	return dp[n]
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
