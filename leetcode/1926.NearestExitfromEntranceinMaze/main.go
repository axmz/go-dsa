package main

import "fmt"

type move struct {
	coord []int
	step  int
}

const (
	wall = '+'
)

var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func nearestExit(maze [][]byte, entrance []int) int {
	cols := len(maze[0]) - 1
	rows := len(maze) - 1
	q := []move{{coord: entrance, step: 0}}
	maze[entrance[0]][entrance[1]] = wall // mark entrance as visited

	for len(q) > 0 {
		pop := q[0]
		q = q[1:]

		for _, direction := range directions {
			newRow := pop.coord[0] + direction[0]
			newCol := pop.coord[1] + direction[1]

			// Check if it's a valid move
			if newRow < 0 || newRow > rows || newCol < 0 || newCol > cols || maze[newRow][newCol] == wall {
				continue
			}

			// Check if it's an exit
			if (newRow == 0 || newRow == rows || newCol == 0 || newCol == cols) && !(newRow == entrance[0] && newCol == entrance[1]) {
				return pop.step + 1
			}

			// Mark as visited before adding to queue and add to queue, otherwise TLE, MLE
			maze[newRow][newCol] = wall
			q = append(q, move{coord: []int{newRow, newCol}, step: pop.step + 1})
		}
	}

	return -1
}

func main() {
	// maze := [][]byte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}}
	// entrance := []int{1, 2}
	// maze := [][]byte{{'+', '+', '+'}, {'.', '.', '.'}, {'+', '+', '+'}}
	// entrance := []int{1, 0}
	maze := [][]byte{{'.', '+'}}
	entrance := []int{0, 0}
	res := nearestExit(maze, entrance)
	fmt.Println(res)
}
