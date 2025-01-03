package main

import "fmt"

func wordBreak2(s string, wordDict []string) bool {
	d := make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		d[v] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if ok := d[s[j:i]]; ok && dp[j] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func wordBreak3(s string, wordDict []string) bool {
	dp := make([]bool, len(s))
	for i := range s {
		for _, w := range wordDict {
			if i < len(w)-1 {
				continue
			}
			if i == len(w)-1 || dp[i-len(w)] {
				if s[i-len(w)+1:i+1] == w {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[len(s)-1]
}

func wordBreak(s string, wordDict []string) bool {
	memo := make(map[int]bool) // Map to store results for memoization

	var dp func(i int) bool

	// Define the recursive function
	dp = func(i int) bool {
		if i == len(s) {
			return true
		}

		// Return cached result if already computed
		if res, exists := memo[i]; exists {
			return res
		}

		for _, w := range wordDict {
			j := i + len(w)

			// Ensure we don't go out of bounds and check substring
			if j <= len(s) && s[i:j] == w {
				if dp(j) {
					memo[i] = true
					return true
				}
			}
		}

		// Store the result as false if no valid segmentation is found
		memo[i] = false
		return false
	}

	// Start recursion from index 0
	return dp(0)
}

func main() {
	s := "bb"
	wordDict := []string{"a", "b", "bbb", "bbbb"}
	// s := "applepenapple"
	// wordDict := []string{"apple", "pen"}
	// s := "catsandog"
	// wordDict := []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s, wordDict))
}
