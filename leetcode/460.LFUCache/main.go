package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	key       int
	value     int
	freq      int
	timestamp int
	index     int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].freq == pq[j].freq {
		return pq[i].timestamp < pq[j].timestamp
	}
	return pq[i].freq < pq[j].freq
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

func (pq *PriorityQueue) update(item *Item, value int, timestamp int) {
	item.value = value
	item.freq++
	item.timestamp = timestamp
	heap.Fix(pq, item.index)
}

type LFUCache struct {
	cap     int
	pq      PriorityQueue
	m       map[int]*Item
	counter int
}

func Constructor(capacity int) LFUCache {
	pq := make(PriorityQueue, 0, capacity)
	heap.Init(&pq)
	m := make(map[int]*Item)

	return LFUCache{
		cap:     capacity,
		pq:      pq,
		m:       m,
		counter: 0,
	}
}

func (this *LFUCache) Get(key int) int {
	this.counter++
	if item, ok := this.m[key]; ok {
		this.pq.update(item, item.value, this.counter)
		return item.value
	} else {
		return -1
	}
}

func (this *LFUCache) Put(key int, value int) {
	this.counter++
	if item, ok := this.m[key]; ok {
		// find the item in the map
		// if found, update item value and counter++
		this.pq.update(item, value, this.counter)
	} else {
		// if not found check if capacity is full
		if this.pq.Len() == this.cap {
			// pop lfu from pq
			p := heap.Pop(&this.pq).(*Item)
			// remove lfu from map
			delete(this.m, p.key)
		}
		// create new item with counter++
		item := &Item{
			key:       key,
			value:     value,
			freq:      1,
			timestamp: this.counter,
		}
		// add new item to heap and map
		this.m[key] = item
		heap.Push(&this.pq, item)
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
