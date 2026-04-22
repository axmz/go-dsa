package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func furthestBuilding2(heights []int, bricks int, ladders int) int {
	h := &MinHeap{}
	heap.Init(h)

	for i := 1; i < len(heights); i++ {
		diff := heights[i] - heights[i-1]
		if diff > 0 {
			heap.Push(h, diff)
		}

		if h.Len() > ladders {
			bricks -= heap.Pop(h).(int)
		}
		if bricks < 0 {
			return i - 1
		}
	}

	return len(heights) - 1

}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	h := &MaxHeap{}
	heap.Init(h)

	for i := 1; i < len(heights); i++ {
		diff := heights[i] - heights[i-1]
		if diff > 0 {
			heap.Push(h, diff)
			bricks -= diff
		}

		if bricks < 0 {
			if ladders == 0 {
				return i - 1
			}
			bricks += heap.Pop(h).(int)
			ladders--
		}
	}

	return len(heights) - 1
}
func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
