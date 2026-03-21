package main

import (
	"container/heap"
	"fmt"
)

// MinHeap solution
// type Pair struct {
// 	key int
// 	val int
// }
// type MinHeap []Pair

// func (h MinHeap) Len() int           { return len(h) }
// func (h MinHeap) Less(i, j int) bool { return h[i].val < h[j].val }
// func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
// func (h *MinHeap) Push(x any)        { *h = append(*h, x.(Pair)) }
// func (h *MinHeap) Pop() any {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }

// func topKFrequent(nums []int, k int) []int {
// 	m := make(map[int]int)
// 	for _, v := range nums {
// 		m[v]++
// 	}

// 	h := &MinHeap{}
// 	heap.Init(h)

// 	i := 0
// 	for key, val := range m {
// 		p := Pair{key: key, val: val}
// 		if i < k {
// 			heap.Push(h, p)
// 			i++
// 			continue
// 		}
// 		if (*h)[0].val < val {
// 			heap.Pop(h)
// 			heap.Push(h, p)
// 		}
// 	}

// 	res := make([]int, k)
// 	for i, p := range *h {
// 		res[i] = p.key
// 	}

// 	return res
// }

// MaxHeap solution
type Pair struct {
	key int
	val int
}
type MaxHeap []Pair

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].val > h[j].val }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(Pair)) }
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	res := []int{}

	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}
	fmt.Println(freq)

	h := &MaxHeap{}
	heap.Init(h)
	for n, f := range freq {
		p := Pair{n, f}
		heap.Push(h, p)
	}
	fmt.Print(h)

	for i := 0; i < k; i++ {
		pop := heap.Pop(h).(Pair)
		fmt.Println(pop)
		res = append([]int{pop.key}, res...)
	}
	fmt.Println(res)

	return res
}

func main() {
	// nums := []int{1, 1, 1, 2, 2, 3, 3, 3, 3, 4, 5, 6}
	// k := 2
	nums := []int{1}
	k := 1

	r := topKFrequent(nums, k)
	fmt.Println(r)
}
