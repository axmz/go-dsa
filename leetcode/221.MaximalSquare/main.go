package main

import (
	"fmt"
)

func min(ints ...int) int {
	m := ints[0]

	for _, v := range ints {
		if v < m {
			m = v
		}
	}

	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maximalSquare(matrix [][]byte) int {
	maximum := 0
	width, height := len(matrix), len(matrix[0])
	memo := make([][]int, width)

	for i := range memo {
		memo[i] = make([]int, height)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dp func(r, c int) int
	dp = func(r, c int) int {
		if r >= width || c >= height {
			return 0
		}

		if memo[r][c] != -1 {
			return memo[r][c]
		}

		w := dp(r, c+1)
		d := dp(r+1, c+1)
		h := dp(r+1, c)

		if matrix[r][c] == '0' {
			memo[r][c] = 0
		} else {
			memo[r][c] = min(w, d, h) + 1
			maximum = max(maximum, memo[r][c])
		}

		return memo[r][c]
	}

	dp(0, 0)

	return maximum * maximum
}

func main() {
	// matrix := [][]byte{
	// 	{'1', '0', '1', '0', '0'},
	// 	{'1', '0', '1', '1', '1'},
	// 	{'1', '1', '1', '1', '1'},
	// 	{'1', '0', '0', '1', '0'},
	// }

	matrix := [][]byte{
		{'0', '1', '1', '0', '0', '1', '0', '1', '0', '1'},
		{'0', '0', '1', '0', '1', '0', '1', '0', '1', '0'},
		{'1', '0', '0', '0', '0', '1', '0', '1', '1', '0'},
		{'0', '1', '1', '1', '1', '1', '1', '0', '1', '0'},
		{'0', '0', '1', '1', '1', '1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0', '1', '1', '1', '1', '0'},
		{'0', '0', '0', '1', '1', '0', '0', '0', '1', '0'},
		{'1', '1', '0', '1', '1', '0', '0', '1', '1', '1'},
		{'0', '1', '0', '1', '1', '0', '1', '0', '1', '1'},
	}

	// matrix := [][]byte{
	// 	{'1', '1', '1'},
	// 	{'1', '1', '1'},
	// 	{'1', '0', '1'},
	// }

	fmt.Println(maximalSquare(matrix))
}
