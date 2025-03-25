package main

import (
	"container/heap"
	"fmt"
)

type Cell struct {
	row    int
	col    int
	effort int
}

type MinHeap []*Cell

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].effort < h[j].effort }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*Cell)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func minimumEffortPath(heights [][]int) int {
	height := len(heights)
	width := len(heights[0])

	visited := make([][]bool, height)
	for r := range visited {
		visited[r] = make([]bool, width)
	}
	visited[0][0] = true

	h := &MinHeap{}
	heap.Init(h)
	heap.Push(h, &Cell{0, 0, 0})

	directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	neighbors := func(r, c int) [][]int {
		res := [][]int{}

		for _, dir := range directions {
			nr, nc := r, c
			nr += dir[0]
			nc += dir[1]
			if nr >= 0 && nr < height && nc >= 0 && nc < width && !visited[nr][nc] {
				res = append(res, []int{nr, nc})
			}
		}

		return res
	}

	for h.Len() > 0 {
		cell := heap.Pop(h).(*Cell)
		if cell.row == height-1 && cell.col == width-1 {
			return cell.effort
		}

		visited[cell.row][cell.col] = true

		for _, neighbor := range neighbors(cell.row, cell.col) {
			nr := neighbor[0]
			nc := neighbor[1]
			nh := heights[nr][nc]
			heap.Push(h, &Cell{
				row:    nr,
				col:    nc,
				effort: max(cell.effort, abs(nh-heights[cell.row][cell.col])),
			})
		}

	}

	return 0
}

func main() {
	heights := [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}
	fmt.Println(minimumEffortPath(heights))
	heights = [][]int{{1, 10, 6, 7, 9, 10, 4, 9}}
	fmt.Println(minimumEffortPath(heights))
	heights = [][]int{{10, 8}, {10, 8}, {1, 2}, {10, 3}, {1, 3}, {6, 3}, {5, 2}}
	fmt.Println(minimumEffortPath(heights))
}
