package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    []int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
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

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value []int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func kClosest(points [][]int, k int) [][]int {
	pq := make(PriorityQueue, len(points))
	for i, point := range points {
		x, y := point[0], point[1]
		// Euclidean distance from the origin (0, 0) is sqrt(x^2 + y^2)
		priority := x*x + y*y
		pq[i] = &Item{
			value:    point,
			priority: priority,
			index:    i,
		}
	}

	heap.Init(&pq)
	res := make([][]int, 0, k)
	for i := 0; i < k; i++ {
		item := heap.Pop(&pq).(*Item)
		res = append(res, item.value)
	}

	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
