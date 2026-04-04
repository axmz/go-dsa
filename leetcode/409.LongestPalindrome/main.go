package main

import "fmt"

func longestPalindrome(s string) int {
	set := [256]byte{}
	res := 0
	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		if set[c] == 0 {
			set[c] = 1
		} else {
			res += 2
			set[c] = 0
		}
	}

	for _, v := range set {
		if v == 1 {
			res++
			break
		}
	}

	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
