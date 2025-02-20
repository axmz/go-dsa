package main

import "fmt"

func gameOfLife(board [][]int) {
	h := len(board)
	w := len(board[0])
	coords := []int{0, 1, -1}

	countLiveNeighbors := func(r, c int) int {
		liveNeighbors := 0
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if j == 0 && i == 0 {
					continue
				}

				rr := r + coords[i]
				cc := c + coords[j]

				if rr < 0 || rr >= h || cc < 0 || cc >= w {
					continue
				}

				val := board[rr][cc]
				if val == 1 || val == -1 {
					liveNeighbors++
				}
			}
		}

		return liveNeighbors
	}

	modifyCell := func(liveNeighbors, r, c int) {
		switch {
		case board[r][c] == 1 && liveNeighbors < 2:
			board[r][c] = -1 // Underpopulation
		case board[r][c] == 1 && (liveNeighbors == 2 || liveNeighbors == 3):
			board[r][c] = 1 // Lives on
		case board[r][c] == 1 && liveNeighbors > 3:
			board[r][c] = -1 // Overpopulation
		case board[r][c] == 0 && liveNeighbors == 3:
			board[r][c] = 2 // Reproduction
		}
	}

	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			n := countLiveNeighbors(r, c)
			modifyCell(n, r, c)
		}
	}

	// modify back
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			switch board[r][c] {
			case -1:
				board[r][c] = 0
			case 2:
				board[r][c] = 1
			}
		}
	}

}

func main() {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(board)
	fmt.Println(board)
}
