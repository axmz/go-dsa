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

type SmallestInfiniteSet struct {
	nums           []int
	currentInteger int
	h              *MinHeap
	m              map[int]bool
}

func Constructor() SmallestInfiniteSet {
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = i + 1
	}

	h := &MinHeap{}
	heap.Init(h)
	return SmallestInfiniteSet{
		nums:           nums,
		currentInteger: 0,
		h:              h,
		m:              make(map[int]bool),
	}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	var pop int
	if s.h.Len() > 0 {
		pop = heap.Pop(s.h).(int)
		delete(s.m, pop)
	} else {
		pop = s.nums[s.currentInteger]
		s.currentInteger++
	}

	return pop
}

func (s *SmallestInfiniteSet) AddBack(num int) {
	if num < s.currentInteger+1 && !s.m[num] {
		heap.Push(s.h, num)
		s.m[num] = true
	}
}

func main() {
	smallestInfiniteSet := Constructor()
	smallestInfiniteSet.AddBack(2)                 // 2 is already in the set, so no change is made.
	fmt.Println(smallestInfiniteSet.PopSmallest()) // return 1, since 1 is the smallest number, and remove it from the set.
	fmt.Println(smallestInfiniteSet.PopSmallest()) // return 2, and remove it from the set.
	fmt.Println(smallestInfiniteSet.PopSmallest()) // return 3, and remove it from the set.
	smallestInfiniteSet.AddBack(1)                 // 1 is added back to the set.
	fmt.Println(smallestInfiniteSet.PopSmallest()) // return 1, since 1 was added back to the set and
	// is the smallest number, and remove it from the set.
	fmt.Println(smallestInfiniteSet.PopSmallest()) // return 4, and remove it from the set.
	fmt.Println(smallestInfiniteSet.PopSmallest()) // return 5, and remove it from the set.
}
