package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type pair struct {
	n1 int
	n2 int
}

type Pairs []pair

func (h Pairs) Len() int           { return len(h) }
func (h Pairs) Less(i, j int) bool { return h[i].n2 > h[j].n2 }
func (h Pairs) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

type MinHeap Pairs

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].n1 < h[j].n1 }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(pair)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getScore(heap *MinHeap, min int) int {
	var sum int
	for _, v := range *heap {
		sum += v.n1
	}
	return sum * min
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	var maxScore int
	pairs := make(Pairs, len(nums1))

	for i, v := range nums1 {
		pairs[i] = pair{n1: v, n2: nums2[i]}
	}

	sort.Sort(pairs)

	var minHeap = &MinHeap{}
	heap.Init(minHeap)
	for i := 0; i < k; i++ {
		heap.Push(minHeap, pairs[i])
	}

	min := pairs[k-1].n2
	maxScore = getScore(minHeap, min)

	for i := k; i < len(nums2); i++ {
		heap.Pop(minHeap)
		heap.Push(minHeap, pairs[i])
		min = pairs[i].n2
		newScore := getScore(minHeap, min)
		if newScore > maxScore {
			maxScore = newScore
		}
	}

	return int64(maxScore)
}

func main() {
	nums1 := []int{1, 3, 3, 2}
	nums2 := []int{2, 1, 3, 4}
	k := 3
	fmt.Println(maxScore(nums1, nums2, k))
}
