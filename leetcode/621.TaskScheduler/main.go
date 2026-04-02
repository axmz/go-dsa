package main

import (
	"container/heap"
	"fmt"
)

// math solution
func leastInterval2(tasks []byte, n int) int {
	freq := [26]int{}
	for _, t := range tasks {
		freq[t-'A']++
	}

	var maxFreq, sameMaxFreq int
	for _, f := range freq {
		if f > maxFreq {
			maxFreq = f
			sameMaxFreq = 1
		} else if f == maxFreq {
			sameMaxFreq++
		}
	}

	res := (n+1)*(maxFreq-1) + sameMaxFreq
	if res > len(tasks) {
		return res
	} else {
		return len(tasks)
	}
}

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

// simulation solution
func leastInterval(tasks []byte, n int) int {
	freq := [26]int{}
	for _, t := range tasks {
		freq[t-'A']++
	}

	h := &MaxHeap{}
	heap.Init(h)
	for _, f := range freq {
		if f > 0 {
			heap.Push(h, f)
		}
	}
	time := 0
	for h.Len() > 0 {
		var temp []int
		for i := 0; i <= n; i++ {
			if h.Len() > 0 {
				f := heap.Pop(h).(int)
				if f > 1 {
					temp = append(temp, f-1)
				}
			}
			time++
			if h.Len() == 0 && len(temp) == 0 {
				break
			}
		}

		for _, f := range temp {
			heap.Push(h, f)
		}
	}

	return time
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
