package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxArea := 0

	for l < r {
		newMaxArea := min(height[l], height[r]) * (r - l)

		if newMaxArea > maxArea {
			maxArea = newMaxArea
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxArea
}

func main() {
	height := []int{2, 3, 4, 5, 18, 17, 6}
	// height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(height)
	fmt.Println(res)
}
