package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minFallingPathSum(matrix [][]int) int {
	rows := len(matrix)
	cols := rows

	for r := 1; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if c == 0 {
				matrix[r][c] += min(matrix[r-1][c], matrix[r-1][c+1])
			} else if c == cols-1 {
				matrix[r][c] += min(matrix[r-1][c], matrix[r-1][c-1])
			} else {
				matrix[r][c] += min(matrix[r-1][c], min(matrix[r-1][c-1], matrix[r-1][c+1]))
			}
		}
	}

	minPath := matrix[rows-1][0]
	for c := 1; c < cols; c++ {
		if matrix[rows-1][c] < minPath {
			minPath = matrix[rows-1][c]
		}
	}
	return minPath
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
