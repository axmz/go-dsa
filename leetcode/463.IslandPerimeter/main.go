package main

import "fmt"

func islandPerimeter(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	res := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				res += 4
				if r > 0 && grid[r-1][c] == 1 {
					res -= 2
				}
				if c > 0 && grid[r][c-1] == 1 {
					res -= 2
				}
			}
		}
	}

	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
