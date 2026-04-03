package main

import (
	"fmt"
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}
	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		if interval[0] > cur[1] {
			res = append(res, cur)
			cur = interval
		} else {
			cur[0] = min(cur[0], interval[0])
			cur[1] = max(cur[1], interval[1])
		}
	}
	res = append(res, cur)
	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
