package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Event struct {
	room  int
	end   int
	index int
}

type PriorityQueue []*Event

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].end == pq[j].end {
		return pq[i].room < pq[j].room
	}
	return pq[i].end < pq[j].end
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Event)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Event, priority int) {
	item.end = priority
	heap.Fix(pq, item.index)
}

type RoomHeap []int

func (h RoomHeap) Len() int           { return len(h) }
func (h RoomHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h RoomHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *RoomHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *RoomHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func mostBooked(n int, meetings [][]int) int {
	busy := make(PriorityQueue, 0, n)
	heap.Init(&busy)

	available := make(RoomHeap, n)
	for i := 0; i < n; i++ {
		available[i] = i
	}
	heap.Init(&available)

	roomUse := make([]int, n)

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	for _, m := range meetings {
		start, end := m[0], m[1]
		duration := end - start

		for busy.Len() > 0 && busy[0].end <= start {
			freed := heap.Pop(&busy).(*Event)
			heap.Push(&available, freed.room)
		}

		if available.Len() > 0 {
			roomID := heap.Pop(&available).(int)
			heap.Push(&busy, &Event{room: roomID, end: end})
			roomUse[roomID]++
		} else {
			earliest := heap.Pop(&busy).(*Event)
			earliest.end += duration
			heap.Push(&busy, earliest)
			roomUse[earliest.room]++
		}
	}

	mostUsedRoom := 0
	for i, use := range roomUse {
		if use > roomUse[mostUsedRoom] {
			mostUsedRoom = i
		}
	}
	return mostUsedRoom
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
