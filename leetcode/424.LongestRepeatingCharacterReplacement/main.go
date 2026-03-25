package main

import "fmt"

func characterReplacement(s string, k int) int {
	freq := make(map[byte]int)
	maxFreq := 0

	l := 0
	maxLength := 0

	for r := 0; r < len(s); r++ {
		freq[s[r]]++
		if freq[s[r]] > maxFreq {
			maxFreq = freq[s[r]]
		}

		if r-l+1-maxFreq > k {
			freq[s[l]]--
			l++
		}
		if r-l+1 > maxLength {
			maxLength = r - l + 1
		}
	}

	return maxLength
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
