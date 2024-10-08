package main

import "fmt"

// func totalNQueens1(n int) int {
// 	var res int
// 	var r, c, queensOnBoard int

// 	underAttackBoard := make([][]int, n)
// 	for i := 0; i < n; i++ {
// 		r := make([]int, n)
// 		underAttackBoard[i] = r
// 	}

// 	queensBoard := make([][]int, n)
// 	for i := 0; i < n; i++ {
// 		r := make([]int, n)
// 		queensBoard[i] = r
// 	}

// 	printBoard := make([][]any, n)
// 	for i := 0; i < n; i++ {
// 		r := make([]any, n)
// 		printBoard[i] = r
// 	}

// 	// print := func() {
// 	// 	fmt.Println()
// 	// 	for i := 0; i < n; i++ {
// 	// 		for j := 0; j < n; j++ {
// 	// 			if queensBoard[i][j] == 0 {
// 	// 				printBoard[i][j] = underAttackBoard[i][j]
// 	// 			} else {
// 	// 				printBoard[i][j] = "Q"
// 	// 			}
// 	// 		}
// 	// 	}

// 	// 	for _, v := range printBoard {
// 	// 		fmt.Printf("%2v\n", v)
// 	// 	}
// 	// }

// 	isWithinBounds := func(pos [2]int) bool {
// 		r := pos[0]
// 		c := pos[1]
// 		return r >= 0 && r < n && c >= 0 && c < n
// 	}

// 	moves := [][2]int{
// 		{-1, 0},  // up
// 		{1, 0},   // down
// 		{0, -1},  // left
// 		{0, 1},   // right
// 		{-1, -1}, // up left
// 		{-1, 1},  // up right
// 		{1, -1},  // down left
// 		{1, 1},   // down right
// 	}

// 	moveQueen := func(r, c int, sign int) {
// 		var start, next, move [2]int
// 		start = [2]int{r, c}
// 		underAttackBoard[r][c] += sign
// 		queensBoard[r][c] += sign
// 		for _, move = range moves {
// 			next = [2]int{start[0] + move[0], start[1] + move[1]}
// 			for isWithinBounds(next) {
// 				r := next[0]
// 				c := next[1]
// 				underAttackBoard[r][c] += sign
// 				next = [2]int{next[0] + move[0], next[1] + move[1]}
// 			}
// 		}
// 	}

// 	placeQueen := func(r, c int) {
// 		moveQueen(r, c, 1)
// 		queensOnBoard++
// 		if queensOnBoard == n {
// 			res++
// 		}
// 		// print()
// 	}

// 	removeQueen := func(r, c int) {
// 		moveQueen(r, c, -1)
// 		queensOnBoard--
// 		// print()
// 	}

// 	isUnderAttack := func(r, c int) bool {
// 		return underAttackBoard[r][c] > 0
// 	}

// 	nextMove := func(r, c int) (int, int) {
// 		nr, nc := r, c
// 		if c+1 >= n {
// 			nr = r + 1
// 			nc = 0
// 		} else {
// 			nc = c + 1
// 		}

// 		return nr, nc
// 	}

// 	var backtrack func(r, c int)
// 	backtrack = func(r, c int) {
// 		if !isWithinBounds([2]int{r, c}) {
// 			return
// 		}

// 		nr, nc := nextMove(r, c)
// 		if isUnderAttack(r, c) {
// 			backtrack(nr, nc)
// 		} else {
// 			placeQueen(r, c)
// 			backtrack(nr, nc)
// 			removeQueen(r, c)
// 			backtrack(nr, nc)
// 		}
// 	}

// 	backtrack(r, c)
// 	return res
// }

func totalNQueens(n int) int {
	var res int
	var r, c, queensOnBoard int

	board := make([]int, n)
	for i := range board {
		board[i] = -1
	}

	isWithinBounds := func(pos [2]int) bool {
		r := pos[0]
		c := pos[1]
		return r >= 0 && r < n && c >= 0 && c < n
	}

	placeQueen := func(r, c int) {
		board[r] = c
		queensOnBoard++
		if queensOnBoard == n {
			res++
		}
	}

	removeQueen := func(r, c int) {
		board[r] = -1
		queensOnBoard--
	}

	isUnderAttack := func(r, c int) bool {
		for i := 0; i <= r; i++ {
			br, bc := i, board[i]
			if bc == -1 {
				continue
			}

			diagonal, adiagonal := br-bc, bc+br

			if br == r ||
				bc == c ||
				diagonal == r-c ||
				adiagonal == c+r {
				return true
			}
		}

		return false
	}

	nextMove := func(r, c int) (int, int) {
		nr, nc := r, c
		if c+1 >= n {
			nr = r + 1
			nc = 0
		} else {
			nc = c + 1
		}

		return nr, nc
	}

	var backtrack func(r, c int)
	backtrack = func(r, c int) {
		if !isWithinBounds([2]int{r, c}) {
			return
		}

		nr, nc := nextMove(r, c)
		if isUnderAttack(r, c) {
			backtrack(nr, nc)
		} else {
			placeQueen(r, c)
			backtrack(nr, nc)
			removeQueen(r, c)
			backtrack(nr, nc)
		}
	}

	backtrack(r, c)
	return res
}

func main() {
	n := 9
	fmt.Println(totalNQueens(n))
}
