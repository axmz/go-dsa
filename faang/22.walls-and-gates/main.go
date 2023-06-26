package main

import (
	"fmt"
	"math"
)

const (
	WALL  = -1
	GATE  = 0
	EMPTY = math.MaxInt32
	INFN  = EMPTY
)

type coord struct {
	row int
	col int
}

func dfs(grid *[][]int, c coord, step int) {
	width := len((*grid)[0]) - 1
	height := len(*grid) - 1

	isValidPath := func(c coord) bool {
		isOutOfBound := c.row < 0 || c.row > height || c.col < 0 || c.col > width
		if isOutOfBound {
			return false
		}

		path := (*grid)[c.row][c.col]
		if path == WALL || path == GATE {
			return false
		}

		return true
	}

	up := coord{c.row - 1, c.col}
	right := coord{c.row, c.col + 1}
	down := coord{c.row + 1, c.col}
	left := coord{c.row, c.col - 1}

	directions := [4]coord{up, right, down, left}

	for _, dir := range directions {
		if isValidPath(dir) {
			currentCellValue := &(*grid)[dir.row][dir.col]
			if step < *currentCellValue {
				*currentCellValue = step
				dfs(grid, dir, step+1)
			}
		}
	}
}

func wallsAndGates(grid [][]int) {
	for i, r := range grid {
		for j, c := range r {
			if c == GATE {
				dfs(&grid, coord{i, j}, 1)
			}
		}
	}
}

func main() {
	grid := [][]int{
		{INFN, WALL, GATE, INFN},
		{INFN, INFN, INFN, GATE},
		{INFN, WALL, INFN, WALL},
		{GATE, WALL, INFN, INFN},
	}

	wallsAndGates(grid)
	fmt.Println(grid)
	// Should be
	// [3 -1  0  1]
	// [2  2  1  0]
	// [1 -1  2 -1]
	// [0 -1  3  4]
}
