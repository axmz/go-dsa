package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    byte
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority > pq[j].priority }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x any)        { *pq = append(*pq, x.(*Item)) }
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return item
}

func reorganizeString(s string) string {
	var count [26]int
	for _, c := range s {
		count[c-'a']++
	}

	pq := make(PriorityQueue, 0, 26)
	for i, v := range count {
		if v > 0 {
			heap.Push(&pq, &Item{value: byte('a' + i), priority: v})
		}
	}

	result := make([]byte, 0, len(s))
	var prev *Item
	for pq.Len() > 0 {
		if prev != nil {
			// clever. wait for the next item then put prev back
			heap.Push(&pq, prev)
			prev = nil
		}
		curr := heap.Pop(&pq).(*Item)
		result = append(result, curr.value)
		curr.priority--
		if curr.priority > 0 {
			prev = curr
		}
	}

	if len(result) != len(s) {
		return ""
	}
	return string(result)
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
