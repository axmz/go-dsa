package main

import "fmt"

func pacificAtlantic(heights [][]int) [][]int {
	rows := len(heights)
	cols := len(heights[0])
	dirrections := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	marked_pacific := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		marked_pacific[i] = make([]bool, cols)
	}

	// mark all pacific
	// can be combined with atlantic
	queue := make([][2]int, 0)
	for r := range heights {
		queue = append(queue, [2]int{r, 0})
	}
	for c := range heights[0] {
		queue = append(queue, [2]int{0, c})
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		r, c := cur[0], cur[1]
		marked_pacific[r][c] = true // mark as pacific reachable
		for _, d := range dirrections {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && !marked_pacific[nr][nc] && heights[nr][nc] >= heights[r][c] {
				queue = append(queue, [2]int{nr, nc})
			}
		}
	}

	// mark all atlantic
	marked_atlantic := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		marked_atlantic[i] = make([]bool, cols)
	}

	queue = make([][2]int, 0)
	for r := range heights {
		queue = append(queue, [2]int{r, cols - 1})
	}
	for c := range heights[0] {
		queue = append(queue, [2]int{rows - 1, c})
	}

	res := make([][]int, 0)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		r, c := cur[0], cur[1]
		marked_atlantic[r][c] = true // mark as atlantic reachable
		for _, d := range dirrections {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && !marked_atlantic[nr][nc] && heights[nr][nc] >= heights[r][c] {
				queue = append(queue, [2]int{nr, nc})
			}
		}
	}

	for r := range heights {
		for c := range heights[0] {
			if marked_pacific[r][c] && marked_atlantic[r][c] {
				res = append(res, []int{r, c})
			}
		}
	}

	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
