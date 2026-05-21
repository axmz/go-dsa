package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxTurbulenceSize(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	maxLen := 1
	up, down := 1, 1

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			up = down + 1
			down = 1
		} else if arr[i] < arr[i-1] {
			down = up + 1
			up = 1
		} else {
			up, down = 1, 1
		}
		maxLen = max(maxLen, max(up, down))
	}

	return maxLen
}

func main() {
	nums := []int{9, 4, 2, 10, 7, 8, 8, 1, 9}
	x := maxTurbulenceSize(nums)
	fmt.Println(x, nums)
}
