package main

import (
	"fmt"
)

var (
	dirrections = [][]int{
		{-2, -1},
		{-2, 1},
		{-1, -2},
		{-1, 2},
		{1, 2},
		{1, -2},
		{2, 1},
		{2, -1},
	}
)

func recursion(n int, k int, row int, column int, memo map[int][][]float64) float64 {

	isInBound := row < n && row >= 0 && column < n && column >= 0

	if !isInBound {
		return 0
	}

	if k == 0 {
		return 1
	}

	if prob := memo[k][row][column]; prob != -1 {
		return prob
	}

	res := 0.00
	for _, v := range dirrections {
		res += recursion(n, k-1, row+v[0], column+v[1], memo) / 8
	}

	memo[k][row][column] = res
	return res
}

func createBoardsMemo(k, n int) map[int][][]float64 {
	var memo = make(map[int][][]float64, k)
	for K := 1; K < k+1; K++ {
		var board = make([][]float64, n)
		for col := 0; col < n; col++ {
			newRow := make([]float64, n)
			for i := range newRow {
				newRow[i] = -1
			}
			board[col] = newRow
		}

		memo[K] = board
	}
	return memo
}

func knightProbability(n int, k int, row int, column int) float64 {
	var probability float64 = 0
	var memo = createBoardsMemo(k, n)

	probability += recursion(n, k, row, column, memo)
	return probability
}

func main() {
	// var n int = 3
	// var k = 2
	// var row = 0
	// var column = 0
	var n int = 3
	var k = 1
	var row = 1
	var column = 1
	res := knightProbability(n, k, row, column)
	fmt.Println(res)
}
