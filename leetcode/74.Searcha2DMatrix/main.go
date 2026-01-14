package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	r := len(matrix)
	c := len(matrix[0])
	left := 0
	right := r*c - 1

	for left <= right {
		mid := left + (right-left)/2
		midValue := matrix[mid/c][mid%c]
		if midValue == target {
			return true
		} else if midValue < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
