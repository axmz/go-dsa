package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func getSkyline(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return [][]int{}
	}

	points := make([][2]int, 0, len(buildings)*2)
	for _, b := range buildings {
		left, right, height := b[0], b[1], b[2]
		points = append(points, [2]int{left, -height})
		points = append(points, [2]int{right, height})
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i][0] != points[j][0] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] < points[j][1]
	})

	skyline := [][]int{}
	maxHeap := &IntHeap{}
	heap.Init(maxHeap)

	active := make(map[int]int)
	heap.Push(maxHeap, 0)
	active[0] = 1

	prevHeight := 0

	for _, point := range points {
		x, h := point[0], point[1]

		if h < 0 {
			height := -h
			active[height]++
			heap.Push(maxHeap, height)
		} else {
			height := h
			active[height]--
		}

		for maxHeap.Len() > 0 && active[(*maxHeap)[0]] == 0 {
			heap.Pop(maxHeap)
		}

		currentHeight := (*maxHeap)[0]

		if currentHeight != prevHeight {
			skyline = append(skyline, []int{x, currentHeight})
			prevHeight = currentHeight
		}
	}

	return skyline
}

// Max Heap
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func main() {
	// nums := [][]int{{0, 2, 3}, {2, 5, 3}}
	nums := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}
	fmt.Println(getSkyline(nums))
}
