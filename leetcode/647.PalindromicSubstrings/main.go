package main

import "fmt"

func countSubstrings(s string) int {
	count := 0
	expandAroundCenter := func(a, b int) {
		for i := 0; a-i >= 0 && b+i < len(s) && s[a] == s[b] && s[a-i] == s[b+i]; i++ {
			count++
		}
	}

	for i := 0; i < len(s); i++ {
		expandAroundCenter(i, i)
		expandAroundCenter(i, i+1)
	}
	return count
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
