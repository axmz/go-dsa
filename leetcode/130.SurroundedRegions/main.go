package main

import "fmt"

func solve2(board [][]byte) {
	h := len(board)
	w := len(board[0])

	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}

	dirrections := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			if board[r][c] == 'X' || visited[r][c] {
				continue
			}

			var contaminated bool
			zeroArea := make([][]int, 0)
			q := [][]int{{r, c}}
			visited[r][c] = true

			for len(q) > 0 {
				pop := q[0]
				q = q[1:]
				zeroArea = append(zeroArea, pop)

				rr, cc := pop[0], pop[1]

				if rr == 0 || rr == h-1 || cc == 0 || cc == w-1 {
					contaminated = true
				}
				for _, dir := range dirrections {
					nr := rr + dir[0]
					nc := cc + dir[1]
					if nr >= 0 &&
						nr < h &&
						nc >= 0 &&
						nc < w &&
						!visited[nr][nc] &&
						board[nr][nc] != 'X' {
						visited[nr][nc] = true
						zeroArea = append(zeroArea, []int{nr, nc})
						q = append(q, []int{nr, nc})
						if nr == 0 || nr == h-1 || nc == 0 || nc == w-1 {
							contaminated = true
						}
					}
				}

			}
			if !contaminated {
				for _, coord := range zeroArea {
					board[coord[0]][coord[1]] = 'X'
				}
			}
		}
	}
}

func solve(board [][]byte) {
	h := len(board)
	w := len(board[0])

	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}

	dirrections := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	q := make([][]int, 0)

	for r := 0; r < h; r++ {
		q = append(q, []int{r, 0}, []int{r, w - 1})
		visited[r][0] = true
		visited[r][w-1] = true
	}

	for c := 0; c < w; c++ {
		q = append(q, []int{0, c}, []int{h - 1, c})
		visited[0][c] = true
		visited[h-1][c] = true
	}

	for len(q) > 0 {
		pop := q[0]
		q = q[1:]
		r := pop[0]
		c := pop[1]
		cell := board[r][c]

		if cell == 'O' {
			board[r][c] = 'E'
			for _, dir := range dirrections {
				nr, nc := r+dir[0], c+dir[1]
				if nr >= 0 && nr < h && nc >= 0 && nc < w && !visited[nr][nc] && board[nr][nc] == 'O' {
					q = append(q, []int{nr, nc})
				}
			}
		}
	}

	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			cell := board[r][c]
			if cell == 'O' {
				board[r][c] = 'X'
			} else if cell == 'E' {
				board[r][c] = 'O'
			}
		}
	}
}

func main() {
	// board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	// board := [][]byte{{'O', 'X', 'X', 'O', 'X'}, {'X', 'O', 'O', 'X', 'O'}, {'X', 'O', 'X', 'O', 'X'}, {'O', 'X', 'O', 'O', 'O'}, {'X', 'X', 'O', 'X', 'O'}}
	board := [][]byte{{'O', 'O', 'O', 'O', 'O', 'O'}, {'O', 'X', 'X', 'X', 'X', 'O'}, {'O', 'X', 'O', 'O', 'X', 'O'}, {'O', 'X', 'O', 'O', 'X', 'O'}, {'O', 'X', 'X', 'X', 'X', 'O'}, {'O', 'O', 'O', 'O', 'O', 'O'}}
	fmt.Println(board)
	solve(board)
	fmt.Println(board)
}
