package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Item struct {
	project  int
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

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	h := &PriorityQueue{}
	heap.Init(h)
	projects := make([][2]int, len(profits))
	for i := 0; i < len(profits); i++ {
		projects[i] = [2]int{profits[i], capital[i]}
	}

	// sort projects by capital required
	sort.Slice(projects, func(i, j int) bool {
		return projects[i][1] < projects[j][1]
	})

	visited := 0
	for i := 0; i < k; i++ {
		for j := visited; j < len(projects); j++ {
			if projects[j][1] > w {
				break
			}
			heap.Push(h, &Item{project: j, priority: projects[j][0]})
			visited++
		}

		if h.Len() == 0 {
			break
		}

		item := heap.Pop(h).(*Item)
		w += item.priority
	}

	return w
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
