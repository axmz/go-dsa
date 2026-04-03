package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minMeetingRooms2(intervals [][]int) int {
	h := &MinHeap{}
	heap.Init(h)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// We can either free all the rooms before pushing a new one, or
	// maxRooms := 0
	// for _, interval := range intervals {
	// 	start, end := interval[0], interval[1]
	// 	for h.Len() > 0 && (*h)[0] <= start {
	// 		heap.Pop(h)
	// 	}
	// 	heap.Push(h, end)
	// 	if h.Len() > maxRooms {
	// 		maxRooms = h.Len()
	// 	}
	// }

	// return maxRooms

	// We can eliminate only one room at a time.
	for _, interval := range intervals {
		if h.Len() > 0 && (*h)[0] <= interval[0] {
			heap.Pop(h)
		}
		heap.Push(h, interval[1])
	}
	return h.Len()
}

// Nice solution
func minMeetingRooms(intervals [][]int) int {
	starts := make([]int, len(intervals))
	ends := make([]int, len(intervals))
	for i, interval := range intervals {
		starts[i] = interval[0]
		ends[i] = interval[1]
	}

	sort.Ints(starts)
	sort.Ints(ends)

	rooms := 0
	maxRooms := 0
	for i, j := 0, 0; i < len(starts); {
		if starts[i] < ends[j] {
			i++
			rooms++
		} else {
			j++
			rooms--
		}

		if rooms > maxRooms {
			maxRooms = rooms
		}

	}

	return maxRooms
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
