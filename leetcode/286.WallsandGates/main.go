package main

import "fmt"

var directions = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func wallsAndGates(rooms [][]int) {
	rows := len(rooms)
	cols := len(rooms[0])
	q := [][2]int{}
	// Multi-Source BFS: enqueue all gates first
	// Reduces memory consumption and allocs
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if rooms[r][c] == 0 {
				q = append(q, [2]int{r, c})
			}
		}
	}

	for len(q) > 0 {
		pop := q[0]
		q = q[1:]
		r, c := pop[0], pop[1]

		for _, dir := range directions {
			newR := r + dir[0]
			newC := c + dir[1]
			// nice technique to split a long if condition
			withinBounds := newR >= 0 && newR < rows && newC >= 0 && newC < cols
			if withinBounds {
				// has check withing bounds first before accessing rooms[newR][newC]
				notWallOrGate := rooms[newR][newC] != -1 && rooms[newR][newC] != 0
				if notWallOrGate {
					// split a long if condition
					shorterPath := rooms[r][c]+1 < rooms[newR][newC]
					if shorterPath {
						rooms[newR][newC] = rooms[r][c] + 1 // this technique allows me to avoid passing distance in the queue ex: [r,c,dist]
						q = append(q, [2]int{newR, newC})   // if we don't use Multi-Source BFS, we get MLE here
					}
				}
			}
		}
	}
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
