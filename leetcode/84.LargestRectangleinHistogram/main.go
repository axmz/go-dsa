package main

import "fmt"

// func calculateArea(heights []int, start, end int) int {
// 	if start > end {
// 		return 0
// 	}
// 	minIndex := start
// 	for i := start; i <= end; i++ {
// 		if heights[i] < heights[minIndex] {
// 			minIndex = i
// 		}
// 	}

// 	currentArea := heights[minIndex] * (end - start + 1)
// 	leftArea := calculateArea(heights, start, minIndex-1)
// 	rightArea := calculateArea(heights, minIndex+1, end)

// 	return max(currentArea, max(leftArea, rightArea))
// }

// func largestRectangleArea(heights []int) int {
// 	return calculateArea(heights, 0, len(heights)-1)

// }

func largestRectangleArea(heights []int) int {
	stack := []int{}
	maxArea := 0
	heights = append(heights, 0) // Append a sentinel value to pop all elements at the end

	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			var w int
			if len(stack) == 0 {
				w = i
			} else {
				w = i - stack[len(stack)-1] - 1
			}
			maxArea = max(maxArea, h*w)
		}
		stack = append(stack, i)
	}
	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 9, 1, 2}
	fmt.Println(largestRectangleArea(nums))
}
