package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	for i := 1; i < rows; i++ {
		grid[i][0] += grid[i-1][0]
	}

	for j := 1; j < cols; j++ {
		grid[0][j] += grid[0][j-1]
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[rows-1][cols-1]
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
