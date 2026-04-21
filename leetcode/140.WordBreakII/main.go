package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) []string {
	l := len(s)
	memo := make(map[int][]string)

	words := make(map[string]struct{})
	for _, w := range wordDict {
		words[w] = struct{}{}
	}

	var backtrack func(start int) []string
	backtrack = func(start int) []string {
		if start >= l {
			return []string{""}
		}

		if val, exist := memo[start]; exist {
			return val
		}

		res := []string{}

		for i := start; i < l; i++ {
			sub := s[start : i+1]
			if _, exist := words[sub]; exist {
				next := backtrack(i + 1)
				for _, n := range next {
					if n == "" {
						res = append(res, sub)
					} else {
						res = append(res, sub+" "+n)
					}
				}
			}
		}

		memo[start] = res
		return res
	}

	return backtrack(0)
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
