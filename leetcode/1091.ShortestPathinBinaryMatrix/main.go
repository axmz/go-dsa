package main

import "fmt"

func shortestPathBinaryMatrix(grid [][]int) int {
	h := len(grid)
	w := len(grid[0])
	if grid[0][0] == 1 || grid[h-1][w-1] == 1 {
		return -1
	}

	q := [][]int{{0, 0}}
	var pop []int
	grid[0][0] = 1

	getNeighbours := func(r, c int) [][]int {
		res := [][]int{}
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				newR := r + i
				newC := c + j
				if newR >= 0 && newR < h && newC >= 0 && newC < w && grid[newR][newC] == 0 {
					res = append(res, []int{newR, newC})
				}
			}
		}

		return res
	}

	for len(q) > 0 {
		pop, q = q[0], q[1:]
		r := pop[0]
		c := pop[1]
		cell := grid[r][c]
		if cell == grid[h-1][w-1] {
			return cell
		}

		neighbours := getNeighbours(r, c)
		for _, n := range neighbours {
			grid[n[0]][n[1]] = cell + 1
		}

		q = append(q, neighbours...)
	}

	return -1
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
