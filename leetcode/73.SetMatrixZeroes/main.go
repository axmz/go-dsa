package main

import "fmt"

func setZeroes(matrix [][]int) {
	h := len(matrix)
	w := len(matrix[0])

	firstRow := false
	for i := range matrix[0] {
		if matrix[0][i] == 0 {
			firstRow = true
			break
		}
	}

	firstCol := false
	for j := 0; j < h; j++ {
		if matrix[j][0] == 0 {
			firstCol = true
			break
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := h - 1; i > 0; i-- {
		if matrix[i][0] == 0 {
			for ii := range matrix[i] {
				matrix[i][ii] = 0
			}
		}
	}

	for j := w - 1; j > 0; j-- {
		if matrix[0][j] == 0 {
			for jj := 0; jj < h; jj++ {
				matrix[jj][j] = 0
			}
		}
	}

	if firstRow {
		for ii := range matrix[0] {
			matrix[0][ii] = 0
		}
	}

	if firstCol {
		for jj := 0; jj < h; jj++ {
			matrix[jj][0] = 0
		}
	}
}

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	fmt.Println(matrix)
}
