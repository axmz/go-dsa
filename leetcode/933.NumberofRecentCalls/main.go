package main

import (
	"fmt"
	"sort"
)

const timeframe = 3000

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return *new(RecentCounter)
}

func (this *RecentCounter) Ping(t int) int {
	this.queue = append(this.queue, t)
	dif := t - timeframe
	idx := sort.Search(len(this.queue), func(i int) bool {
		return this.queue[i] >= dif
	})

	this.queue = this.queue[idx:]
	return len(this.queue)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

func main() {
	obj := Constructor()
	obj.Ping(5)
	obj.Ping(100)
	obj.Ping(2000)
	obj.Ping(3000)
	last := obj.Ping(6006)
	fmt.Println(last)
}
