package main

import "fmt"

func rotate(matrix [][]int) {
	l := len(matrix)
	for i := 0; i < (l+1)/2; i++ {
		for j := 0; j < l/2; j++ {
			n := l - 1
			temp := matrix[n-j][i]
			matrix[n-j][i] = matrix[n-i][n-j]
			matrix[n-i][n-j] = matrix[j][n-i]
			matrix[j][n-i] = matrix[i][j]
			matrix[i][j] = temp
		}
	}
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	fmt.Println(matrix)
}
