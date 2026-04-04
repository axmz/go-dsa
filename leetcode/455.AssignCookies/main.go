package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	count := 0
	for i, j := 0, 0; i < len(g) && j < len(s); {
		if g[i] <= s[j] {
			count++
			i++
		}
		j++
	}
	return count
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
