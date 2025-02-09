package main

import (
	"fmt"
	"math"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	k := math.MinInt
	ans := 0

	for _, v := range points {
		x, y := v[0], v[1]

		if x > k {
			k = y
			ans++
		}
	}

	return ans
}

func main() {
	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	fmt.Println(findMinArrowShots(points))
}
