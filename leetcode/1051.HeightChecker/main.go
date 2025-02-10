package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	count := 0
	sorted := make([]int, len(heights))
	copy(sorted, heights)
	sort.Ints(sorted)
	for i := range heights {
		if heights[i] != sorted[i] {
			count++
		}
	}
	return count
}

func main() {
	nums := []int{1, 1, 4, 2, 1, 3}
	fmt.Println(heightChecker(nums))
}
