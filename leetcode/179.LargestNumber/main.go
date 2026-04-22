package main

import (
	"fmt"
	"sort"
	"strings"
)

func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = fmt.Sprintf("%d", num)
	}
	sort.Slice(strs, func(a, b int) bool {
		s1 := strs[a] + strs[b]
		s2 := strs[b] + strs[a]
		return s1 > s2
	})

	if strs[0] == "0" {
		return "0"
	}

	return strings.Join(strs, "")
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
