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

func totalCost(costs []int, k int, candidates int) int64 {
	var leftTail, rightHead int

	if k > len(costs) {
		k = len(costs)
	}

	if len(costs) >= 2*candidates {
		leftTail = candidates - 1
		rightHead = len(costs) - candidates
	} else {
		leftTail = len(costs)/2 - 1
		rightHead = leftTail + 1
	}

	lh := &MinHeap{}
	heap.Init(lh)
	for i := 0; i <= leftTail; i++ {
		heap.Push(lh, costs[i])
	}

	rh := &MinHeap{}
	heap.Init(rh)
	for i := rightHead; i < len(costs); i++ {
		heap.Push(rh, costs[i])
	}

	var sum int
	for i := 0; i < k; i++ {
		if len(*lh) == 0 && len(*rh) == 0 {
			break
		}

		if len(*rh) == 0 || (len(*lh) > 0 && (*lh)[0] <= (*rh)[0]) {
			sum += heap.Pop(lh).(int)
			if leftTail+1 < rightHead {
				leftTail++
				heap.Push(lh, costs[leftTail])
			}
		} else {
			sum += heap.Pop(rh).(int)
			if rightHead > leftTail+1 {
				rightHead--
				heap.Push(rh, costs[rightHead])
			}
		}
	}

	return int64(sum)
}

func main() {
	// costs := []int{17, 12, 10, 2, 7, 2, 11, 20, 8}
	// k := 3
	// candidates := 4

	// costs := []int{1, 2, 4, 1}
	// k := 3
	// candidates := 3

	// costs := []int{31, 25, 72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 17, 58}
	// k := 11
	// candidates := 2

	costs := []int{57, 33, 26, 76, 14, 67, 24, 90, 72, 37, 30}
	k := 11
	candidates := 2

	fmt.Println(totalCost(costs, k, candidates))
}
