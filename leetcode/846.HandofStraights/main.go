package main

import (
	"container/heap"
	"fmt"
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

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	h := &MinHeap{}
	heap.Init(h)
	freq := make(map[int]int)
	for _, card := range hand {
		if freq[card] == 0 {
			heap.Push(h, card)
		}
		freq[card]++
	}

	for h.Len() > 0 {
		card := (*h)[0]
		for i := 0; i < groupSize; i++ {
			c := card + i
			if freq[c] == 0 {
				return false
			}
			freq[c]--
			if freq[c] == 0 {
				delete(freq, c)
				if c != (*h)[0] {
					return false
				}
				heap.Pop(h)
			}
		}
	}

	return true
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
