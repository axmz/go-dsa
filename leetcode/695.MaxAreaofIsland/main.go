package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	rows := len(grid)
	cols := len(grid[0])
	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if grid[r][c] == 0 {
			return 0
		}

		grid[r][c] = 0
		area := 1
		for _, dir := range dirs {
			newR, newC := dir[0]+r, dir[1]+c
			if newR >= 0 && newR < rows && newC >= 0 && newC < cols {
				area += dfs(newR, newC)
			}
		}

		return area
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			area := dfs(r, c)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
