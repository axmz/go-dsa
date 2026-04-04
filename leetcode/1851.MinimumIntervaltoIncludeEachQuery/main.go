package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type HeapItem struct {
	size int
	end  int
}

type MinHeap []HeapItem

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].size < h[j].size }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(HeapItem)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minInterval(intervals [][]int, queries []int) []int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	type Query struct {
		idx   int
		value int
	}

	q := make([]Query, len(queries))
	for i, query := range queries {
		q[i] = Query{idx: i, value: query}
	}
	sort.Slice(q, func(i, j int) bool {
		return q[i].value < q[j].value
	})

	h := &MinHeap{}
	heap.Init(h)

	result := make([]int, len(queries))
	intervalIdx := 0

	for _, query := range q {
		for intervalIdx < len(intervals) && intervals[intervalIdx][0] <= query.value {
			start := intervals[intervalIdx][0]
			end := intervals[intervalIdx][1]
			heap.Push(h, HeapItem{size: end - start + 1, end: end})
			intervalIdx++
		}

		for h.Len() > 0 && (*h)[0].end < query.value {
			heap.Pop(h)
		}

		if h.Len() > 0 {
			result[query.idx] = (*h)[0].size
		} else {
			result[query.idx] = -1
		}
	}

	return result
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
