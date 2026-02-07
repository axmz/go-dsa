package main

import "fmt"

// Can be implemented in two ways:
// 1. Using a count variable to track the number of elements
// 2. Using an extra space in the array to differentiate between full and empty states

// Implementation using count variable

// type MyCircularQueue struct {
// 	q     []int
// 	head  int
// 	tail  int
// 	size  int
// 	count int
// }

// func Constructor(k int) MyCircularQueue {
// 	return MyCircularQueue{
// 		q:     make([]int, k),
// 		head:  0,
// 		tail:  0,
// 		size:  k,
// 		count: 0,
// 	}
// }

// func (this *MyCircularQueue) EnQueue(value int) bool {
// 	if this.IsFull() {
// 		return false
// 	}
// 	this.q[this.tail] = value
// 	this.tail = (this.tail + 1) % this.size
// 	this.count++
// 	return true
// }

// func (this *MyCircularQueue) DeQueue() bool {
// 	if this.IsEmpty() {
// 		return false
// 	}
// 	this.head = (this.head + 1) % this.size
// 	this.count--
// 	return true
// }

// func (this *MyCircularQueue) Front() int {
// 	if this.IsEmpty() {
// 		return -1
// 	}
// 	return this.q[this.head]
// }

// func (this *MyCircularQueue) Rear() int {
// 	if this.IsEmpty() {
// 		return -1
// 	}
// 	return this.q[(this.tail-1+this.size)%this.size]
// }

// func (this *MyCircularQueue) IsEmpty() bool {
// 	return this.count == 0
// }

// func (this *MyCircularQueue) IsFull() bool {
// 	return this.count == this.size
// }

// Implementation using extra space
type MyCircularQueue struct {
	q    []int
	head int
	tail int
	size int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		q:    make([]int, k+1),
		head: 0,
		tail: 0,
		size: k + 1,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.q[this.tail] = value
	this.tail = (this.tail + 1) % this.size
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % this.size
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.q[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.q[(this.tail-1+this.size)%this.size]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == this.tail
}

func (this *MyCircularQueue) IsFull() bool {
	return (this.tail+1)%this.size == this.head
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

func main() {
	fmt.Println(-1 % 4)
}
