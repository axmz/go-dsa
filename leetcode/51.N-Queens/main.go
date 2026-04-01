package main

import "fmt"

type Board struct {
	size   int
	queens map[[2]int]bool
}

func NewBoard(boardSize int) Board {
	return Board{
		size:   boardSize,
		queens: make(map[[2]int]bool, 0),
	}
}

func (b *Board) Mark(row, col int) {
	b.queens[[2]int{row, col}] = true
}

func (b *Board) Unmark(row, col int) {
	delete(b.queens, [2]int{row, col})
}

func (b *Board) Check(row, col int) bool {
	for queen := range b.queens {
		qr, qc := queen[0], queen[1]
		if qr == row || qc == col {
			return false
		}

		rowDiff := qr - row
		if rowDiff < 0 {
			rowDiff = -rowDiff
		}

		colDiff := qc - col
		if colDiff < 0 {
			colDiff = -colDiff
		}

		if rowDiff == colDiff {
			return false
		}
	}

	return true
}

func (b *Board) Print() []string {
	board := make([]string, b.size)
	for i := 0; i < b.size; i++ {
		row := make([]byte, b.size)
		for j := 0; j < b.size; j++ {
			row[j] = '.'
		}
		board[i] = string(row)
	}

	for queen := range b.queens {
		r, c := queen[0], queen[1]
		row := []byte(board[r])
		row[c] = 'Q'
		board[r] = string(row)
	}

	return board
}

func solveNQueens(n int) [][]string {
	var res [][]string
	board := NewBoard(n)

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			res = append(res, board.Print())
			return
		}

		for col := 0; col < n; col++ {
			if !board.Check(row, col) {
				continue
			}

			board.Mark(row, col)
			backtrack(row + 1)
			board.Unmark(row, col)
		}
	}

	backtrack(0)
	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
