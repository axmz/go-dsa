// the requirement is to only keep top k largest elements in the heap
// for that you need a min-heap
// so that you can easily remove the smallest element when the size exceeds k
package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	topElements *MinHeap
	limit       int
}

func Constructor(k int, nums []int) KthLargest {
	h := &MinHeap{}
	heap.Init(h)
	kl := KthLargest{topElements: h, limit: k}
	for _, num := range nums {
		kl.Add(num)
	}
	return kl
}

func (this *KthLargest) Add(val int) int {
	if this.topElements.Len() < this.limit {
		heap.Push(this.topElements, val)
	} else if val > (*this.topElements)[0] {
		heap.Pop(this.topElements)
		heap.Push(this.topElements, val)
	}

	if this.topElements.Len() > 0 {
		return (*this.topElements)[0]
	}
	return 0
}

func main() {
	h := &MinHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
