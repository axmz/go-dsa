package main

import "fmt"

// Another approach is to use 1 and -1 for player1 and player2
// The player who reaches n or -n first wins.

type TicTacToe struct {
	boardSize          int
	playerRows         [2][]int
	playerCols         [2][]int
	playerDiagonal     [2]int
	playerAntidiagonal [2]int
}

func Constructor(n int) TicTacToe {
	playerRows := [2][]int{}
	playerCols := [2][]int{}
	for i := range playerRows {
		playerRows[i] = make([]int, n)
		playerCols[i] = make([]int, n)
	}

	return TicTacToe{
		boardSize:          n,
		playerRows:         playerRows,
		playerCols:         playerCols,
		playerDiagonal:     [2]int{},
		playerAntidiagonal: [2]int{},
	}
}

func (t *TicTacToe) Move(row, col, player int) int {
	playerIdx := player - 1

	t.playerRows[playerIdx][row]++
	t.playerCols[playerIdx][col]++

	if row == col {
		t.playerDiagonal[playerIdx]++
	}

	if row+col == t.boardSize-1 {
		t.playerAntidiagonal[playerIdx]++
	}

	if t.playerRows[playerIdx][row] == t.boardSize ||
		t.playerCols[playerIdx][col] == t.boardSize ||
		t.playerDiagonal[playerIdx] == t.boardSize ||
		t.playerAntidiagonal[playerIdx] == t.boardSize {
		return player
	}

	return 0
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
