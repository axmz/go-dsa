package main

import (
	"fmt"
	"sort"
)

func canAttendMeetings(intervals [][]int) bool {
	// not needed
	// if len(intervals) < 2 {
	// 	return true
	// }

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < len(intervals); i++ {
		prev := intervals[i-1]
		cur := intervals[i]
		if prev[1] > cur[0] {
			return false
		}
	}

	return true

}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
