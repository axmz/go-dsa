package main

import (
	"container/heap"
	"fmt"
)

type Cell struct {
	row    int
	col    int
	height int
}

type MinHeap []Cell

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].height < h[j].height }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(Cell)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func swimInWater(grid [][]int) int {
	maxWater := 0
	rows := len(grid)
	cols := len(grid[0])
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	visited := make([][]bool, rows)
	for row := 0; row < rows; row++ {
		visited[row] = make([]bool, cols)
	}

	visited[0][0] = true

	h := &MinHeap{}
	heap.Init(h)
	heap.Push(h, Cell{row: 0, col: 0, height: grid[0][0]})

	for h.Len() > 0 {
		cur := heap.Pop(h).(Cell)
		if cur.height > maxWater {
			maxWater = cur.height
		}

		if cur.row == rows-1 && cur.col == cols-1 {
			return maxWater
		}

		for _, d := range directions {
			nr, nc := cur.row+d[0], cur.col+d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
				if !visited[nr][nc] {
					visited[nr][nc] = true
					heap.Push(h, Cell{row: nr, col: nc, height: grid[nr][nc]})
				}
			}
		}
	}

	return -1
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
