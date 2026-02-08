package main

import (
	"fmt"
)

type MinStack struct {
	s    []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{
		s:    make([]int, 0),
		mins: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.s = append(this.s, val)
	if len(this.mins) == 0 {
		this.mins = append(this.mins, val)
	} else if this.mins[len(this.mins)-1] > val {
		this.mins = append(this.mins, val)
	} else {
		this.mins = append(this.mins, this.mins[len(this.mins)-1])
	}
}

func (this *MinStack) Pop() {
	this.s = this.s[:len(this.s)-1]
	this.mins = this.mins[:len(this.mins)-1]
}

func (this *MinStack) Top() int {
	return this.s[len(this.s)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
