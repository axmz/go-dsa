package main

import "fmt"

func findDiagonalOrder(mat [][]int) []int {
	w := len(mat[0])
	h := len(mat)
	res := make([]int, 0, w*h)

	coords := make([][2]int, 0, w+h-1)
	for i := 0; i < w; i++ {
		coords = append(coords, [2]int{i, 0})
		if i == w-1 {
			for j := 1; j < h; j++ {
				coords = append(coords, [2]int{i, j})
			}
		}
	}

	up := [2]int{1, -1}
	down := [2]int{-1, 1}
	getOpposite := func(coord [2]int) [2]int {
		x, y := coord[0], coord[1]
		if x+y < h {
			return [2]int{0, x + y}
		}
		return [2]int{x + y - (h - 1), h - 1}
	}

	op := make([][2]int, len(coords))
	for i, v := range coords {
		op[i] = getOpposite(v)
	}

	for i, coord := range coords {
		var start [2]int
		var direction [2]int
		if i%2 == 0 {
			opposite := getOpposite(coord)
			start = opposite
			direction = up
		} else {
			start = coord
			direction = down
		}

		x, y := start[0], start[1]

		newX, newY := x, y

		for newX >= 0 && newX < w && newY >= 0 && newY < h {
			res = append(res, mat[newY][newX])
			newX, newY = newX+direction[0], newY+direction[1]
		}
	}

	return res
}

func main() {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// mat := [][]int{{7}, {9}, {6}}
	// mat := [][]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(findDiagonalOrder(mat))
}
