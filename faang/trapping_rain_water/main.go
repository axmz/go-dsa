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

func trap(arr []int) int {
	var total int = 0
	l, r := 0, len(arr)-1
	maxL, maxR := arr[l], arr[r]

	var cur int = min(maxL, maxR)

	for l < r {
		cur = min(arr[l], arr[r])
		total += min(maxL, maxR) - cur
		if maxL <= maxR {
			l++
			if arr[l] > maxL {
				maxL = arr[l]
			}
		} else {
			r--
			if arr[r] > maxR {
				maxR = arr[r]
			}
		}
	}

	return total
}

func main() {
	// arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// arr := []int{0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2}
	arr := []int{4, 2, 0, 3, 2, 5}
	res := trap(arr)
	fmt.Printf("Total: %4d", res)
}
