package main

import "fmt"

func canPermutePalindrome(s string) bool {
	l := len(s)
	charCount := make(map[byte]int)
	oddCount := 0
	for i := 0; i < l; i++ {
		charCount[s[i]]++
		if charCount[s[i]]%2 == 1 {
			oddCount++
		} else {
			oddCount--
		}
	}

	return oddCount <= 1
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
