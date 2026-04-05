package main

import "fmt"

func partitionLabels(s string) []int {
	lastIdx := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		lastIdx[s[i]] = i
	}

	res := []int{}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		if lastIdx[s[i]] > end {
			end = lastIdx[s[i]]
		}

		if i == end {
			res = append(res, end-start+1)
			start = end + 1
		}
	}

	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
