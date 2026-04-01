package main

import "fmt"

// with memoization
func partition(s string) [][]string {
	memo := make(map[int][][]string, len(s))

	var dfs func(start int) [][]string
	dfs = func(start int) [][]string {
		if start == len(s) {
			return [][]string{{}}
		}

		if cached, ok := memo[start]; ok {
			return cached
		}

		res := [][]string{}
		for end := start + 1; end <= len(s); end++ {
			sub := s[start:end]
			if !isPalindrome(sub) {
				continue
			}

			for _, suffix := range dfs(end) {
				path := make([]string, 0, len(suffix)+1)
				path = append(path, sub)
				path = append(path, suffix...)
				res = append(res, path)
			}
		}

		memo[start] = res
		return res
	}

	return dfs(0)
}

// no memoization

// func partition(s string) [][]string {
// 	res := [][]string{}
// 	var backtrack func(start int, path []string)
// 	backtrack = func(start int, path []string) {
// 		if start == len(s) {
// 			res = append(res, path)
// 			return
// 		}

// 		for i := start + 1; i <= len(s); i++ {
// 			sub := s[start:i]
// 			if isPalindrome(sub) {
// 				p := append([]string{}, path...)
// 				p = append(p, sub)
// 				backtrack(i, p)
// 			}
// 		}
// 	}
// 	backtrack(0, []string{})
// 	return res
// }

func isPalindrome(s string) bool {
	lo, hi := 0, len(s)-1
	for lo < hi {
		if s[lo] != s[hi] {
			return false
		}
		lo++
		hi--
	}
	return true
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
