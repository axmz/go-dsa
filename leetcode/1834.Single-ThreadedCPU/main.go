package main

import (
	"container/heap"
	"sort"
)

// Two priority queues solution is not optimal
// We need heap when items come and go
// Otherwise just sort the items once
type Task struct {
	idx            int
	processingTime int
}

type ProcessingTimePQ []*Task

func (pq ProcessingTimePQ) Len() int { return len(pq) }

func (pq ProcessingTimePQ) Less(i, j int) bool {
	if pq[i].processingTime == pq[j].processingTime {
		return pq[i].idx < pq[j].idx
	}
	return pq[i].processingTime < pq[j].processingTime
}

func (pq ProcessingTimePQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *ProcessingTimePQ) Push(x any) {
	*pq = append(*pq, x.(*Task))
}

func (pq *ProcessingTimePQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return item
}

func getOrder(tasks [][]int) []int {
	n := len(tasks)
	order := make([]int, n)
	for i := range order {
		order[i] = i
	}
	sort.Slice(order, func(i, j int) bool {
		return tasks[order[i]][0] < tasks[order[j]][0]
	})

	pq := make(ProcessingTimePQ, 0)
	heap.Init(&pq)
	result := make([]int, 0, n)
	currentTime := 0
	i := 0

	for len(result) < n {
		for i < n && tasks[order[i]][0] <= currentTime {
			heap.Push(&pq, &Task{idx: order[i], processingTime: tasks[order[i]][1]})
			i++
		}
		if pq.Len() == 0 {
			currentTime = tasks[order[i]][0]
			continue
		}
		task := heap.Pop(&pq).(*Task)
		result = append(result, task.idx)
		currentTime += task.processingTime
	}

	return result
}

func main() {
}
