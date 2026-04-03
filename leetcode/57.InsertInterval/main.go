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

func insert2(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	res := [][]int{}

	for _, interval := range intervals {
		if interval[1] < newInterval[0] {
			res = append(res, interval)
		} else if interval[0] > newInterval[1] {
			res = append(res, interval)
		} else {
			newInterval[0] = min(newInterval[0], interval[0])
			newInterval[1] = max(newInterval[1], interval[1])
		}
	}

	res = append(res, newInterval)
	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})

	return res
}

func insert(intervals [][]int, newInterval []int) [][]int {
	// If the intervals vector is empty, return a vector containing the newInterval
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	n := len(intervals)
	target := newInterval[0]
	left, right := 0, n-1

	// Binary search to find the position to insert newInterval
	for left <= right {
		mid := (left + right) / 2
		if intervals[mid][0] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// Insert newInterval at the found position
	result := make([][]int, 0)
	result = append(result, intervals[:left]...)
	result = append(result, newInterval)
	result = append(result, intervals[left:]...)

	// Merge overlapping intervals
	res := make([][]int, 0)
	for _, interval := range result {
		// If res is empty or there is no overlap, add the interval to the result
		if len(res) == 0 || res[len(res)-1][1] < interval[0] {
			res = append(res, interval)
			// If there is an overlap, merge the intervals by updating the end of the last interval in res
		} else {
			if res[len(res)-1][1] < interval[1] {
				res[len(res)-1][1] = interval[1]
			}
		}
	}

	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
