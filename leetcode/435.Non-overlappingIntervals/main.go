package main

import (
	"fmt"
	"math"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	k := math.MinInt
	var ans int

	for _, s := range intervals {
		x, y := s[0], s[1]
		if x >= k {
			ans++
			k = y
		}
	}

	return len(intervals) - ans
}

func main() {
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	fmt.Println(eraseOverlapIntervals(intervals))
}
