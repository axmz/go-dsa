package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)

	for _, v := range nums {
		heap.Push(h, v)
	}

	for i := 0; i < k-1; i++ {
		heap.Pop(h)
	}

	return heap.Pop(h).(int)

}

func main() {
	// h := []int{3,2,1,5,6,4}
	// k := 2
	h := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	r := findKthLargest(h, k)
	fmt.Println("res:", r)
}
