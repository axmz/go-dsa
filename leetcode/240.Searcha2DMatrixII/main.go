package main

import "fmt"

// clever. snake method.
func searchMatrix2(matrix [][]int, target int) bool {
	r := len(matrix) - 1
	c := 0
	// starting point can be left-bottom or right-top only
	for r >= 0 && c < len(matrix[0]) {
		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] < target {
			c++
		} else {
			r--
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	var search func(l, r, t, b int) bool
	search = func(l, r, t, b int) bool {
		if l > r || t > b {
			return false
		}
		if target < matrix[t][l] || target > matrix[b][r] {
			return false
		}
		midC := l + (r-l)/2
		midR := t + (b-t)/2
		midValue := matrix[midR][midC]
		if midValue == target {
			return true
		} else if midValue < target {
			return search(midC+1, r, t, b) || search(l, r, midR+1, b)
		} else {
			return search(l, midC-1, t, b) || search(l, r, t, midR-1)
		}
	}

	return search(0, len(matrix[0])-1, 0, len(matrix)-1)
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
