package main

import (
	"fmt"
)

type FreqStack struct {
	freq    map[int]int
	stacks  map[int][]int
	maxFreq int
}

func Constructor() FreqStack {
	return FreqStack{
		freq:   make(map[int]int),
		stacks: make(map[int][]int),
	}
}

func (this *FreqStack) Push(val int) {
	this.freq[val]++
	f := this.freq[val]
	if f > this.maxFreq {
		this.maxFreq = f
	}
	this.stacks[f] = append(this.stacks[f], val)
}

func (this *FreqStack) Pop() int {
	stack := this.stacks[this.maxFreq]
	x := stack[len(stack)-1]
	this.stacks[this.maxFreq] = stack[:len(stack)-1]
	this.freq[x]--
	if len(this.stacks[this.maxFreq]) == 0 {
		this.maxFreq--
	}
	return x
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
