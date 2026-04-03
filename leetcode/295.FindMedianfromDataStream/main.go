package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

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

type MedianFinder struct {
	left  *MaxHeap
	right *MinHeap
}

func Constructor() MedianFinder {
	maxh := &MaxHeap{}
	heap.Init(maxh)
	minh := &MinHeap{}
	heap.Init(minh)
	return MedianFinder{
		left:  maxh,
		right: minh,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.left.Len() == 0 || num <= (*this.left)[0] {
		heap.Push(this.left, num)
	} else {
		heap.Push(this.right, num)
	}

	if this.left.Len() > this.right.Len()+1 {
		x := heap.Pop(this.left).(int)
		heap.Push(this.right, x)
	} else if this.right.Len() > this.left.Len() {
		x := heap.Pop(this.right).(int)
		heap.Push(this.left, x)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() > this.right.Len() {
		return float64((*this.left)[0])
	} else if this.left.Len() < this.right.Len() {
		return float64((*this.right)[0])
	} else {
		return float64((*this.left)[0]+(*this.right)[0]) / 2.0
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
