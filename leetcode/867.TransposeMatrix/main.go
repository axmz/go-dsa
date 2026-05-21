package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	// the size w and h are reversed in the new matrix
	newMatrix := make([][]int, len(matrix[0]))
	for i := range newMatrix {
		newMatrix[i] = make([]int, len(matrix))
	}

	for r := range matrix {
		for c := range matrix[r] {
			newMatrix[c][r] = matrix[r][c]
		}
	}

	return newMatrix
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
