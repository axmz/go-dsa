package main

import (
	"fmt"
)

type MyQueue struct {
	s1 []int
	s2 []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.s1 = append(this.s1, x)
}

func (this *MyQueue) transfer() {
	for len(this.s1) > 0 {
		var x int
		x, this.s1 = this.s1[len(this.s1)-1], this.s1[:len(this.s1)-1]
		this.s2 = append(this.s2, x)
	}
}

func (this *MyQueue) Pop() int {
	if len(this.s2) == 0 {
		this.transfer()
	}

	if len(this.s2) == 0 {
		return 0
	}

	var x int
	x, this.s2 = this.s2[len(this.s2)-1], this.s2[:len(this.s2)-1]
	return x
}

func (this *MyQueue) Peek() int {
	if len(this.s2) == 0 {
		this.transfer()
	}

	if len(this.s2) == 0 {
		return 0
	}

	return this.s2[len(this.s2)-1]
}

func (this *MyQueue) Empty() bool {
	if len(this.s2) == 0 && len(this.s1) == 0 {
		return true
	} else {
		return false
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {

	q := Constructor()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Pop()
	q.Pop()
	q.Push(4)
	q.Push(5)
	// q.Pop()
	// q.Pop()
	r := q.Peek()
	fmt.Println(r)
	fmt.Println(q)
}
