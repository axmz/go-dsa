package main

import (
	"fmt"
)

const (
	FRESH  = 1
	ROTTEN = 2
)

type coord struct {
	row int
	col int
}

func orangesRotting(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	var minutes, fresh int
	var rotten []coord

	for i, r := range grid {
		for j, c := range r {
			if c == FRESH {
				fresh++
			}
			if c == ROTTEN {
				rotten = append(rotten, coord{i, j})
			}
		}
	}

	if fresh == 0 {
		return minutes
	}

	width := len((grid)[0]) - 1
	height := len(grid) - 1
	isFresh := func(c coord) bool {

		isOutOfBound := c.row < 0 || c.row > height || c.col < 0 || c.col > width
		if isOutOfBound {
			return false
		}

		return (grid)[c.row][c.col] == FRESH
	}

	var newRotten []coord
	rot := func(c coord) {
		grid[c.row][c.col] = 2
	}

	for len(rotten) > 0 {
		c := rotten[0]

		up := coord{c.row - 1, c.col}
		down := coord{c.row + 1, c.col}
		left := coord{c.row, c.col - 1}
		right := coord{c.row, c.col + 1}

		directions := [4]coord{up, down, left, right}

		for _, dir := range directions {
			if isFresh(dir) {
				rot(dir)
				fresh--
				newRotten = append(newRotten, dir)
			}
		}

		rotten = rotten[1:]

		if len(rotten) == 0 && len(newRotten) > 0 {
			minutes++
			rotten = append(rotten, newRotten...)
			newRotten = newRotten[:0]
		}
	}

	if fresh > 0 {
		return -1
	}

	return minutes
}

func main() {
	// grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	// grid := [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
	grid := [][]int{
		{0, 2}, // [1] [1]
		{0, 1}, // [1] [1]
		{0, 1}, // [1] [1]
		{1, 1}, // [1] [2]
		{1, 1}, // [2] [2]
		{1, 1}, // [2] [1]
	}
	res := orangesRotting(grid)
	fmt.Println(res)
}
