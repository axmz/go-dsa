package main

import "fmt"

func checkValidString(s string) bool {
	low, high := 0, 0
	// low - tracks the minimum possible number of unmatched (
	// high - tracks the maximum possible number of unmatched (
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			low++
			high++
		case ')':
			low--
			high--
		case '*':
			low--
			high++
		}

		if high < 0 {
			return false
		}

		if low < 0 {
			low = 0
		}
	}

	return low == 0
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
