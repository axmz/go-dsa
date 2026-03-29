package main

import (
	"container/list"
	"fmt"
)

type entry struct {
	key   int
	value int
}

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	order    *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		order:    list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok {
		this.order.MoveToFront(elem)
		return elem.Value.(entry).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}

	if elem, ok := this.cache[key]; ok {
		this.order.MoveToFront(elem)
		elem.Value = entry{key: key, value: value}
	} else {
		if this.order.Len() == this.capacity {
			backElem := this.order.Back()
			this.order.Remove(backElem)
			delete(this.cache, backElem.Value.(entry).key)
		}
		newElem := this.order.PushFront(entry{key: key, value: value})
		this.cache[key] = newElem
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
