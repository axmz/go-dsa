package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    byte
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func longestDiverseString(a int, b int, c int) string {
	pq := make(PriorityQueue, 0, 3)
	if a > 0 {
		heap.Push(&pq, &Item{value: 'a', priority: a})
	}
	if b > 0 {
		heap.Push(&pq, &Item{value: 'b', priority: b})
	}
	if c > 0 {
		heap.Push(&pq, &Item{value: 'c', priority: c})
	}

	var prev *Item
	result := []byte{}

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		limit := 2
		if prev != nil && prev.priority > item.priority {
			limit = 1
		}
		for i := 0; i < limit && item.priority > 0; i++ {
			result = append(result, item.value)
			item.priority--
		}
		if prev != nil {
			heap.Push(&pq, prev)
			prev = nil
		}
		if item.priority > 0 {
			prev = item
		}
	}

	return string(result)
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
