package main

import "fmt"

type Queue struct {
	data []int
}

func (q *Queue) Enqueue(x int) {
	q.data = append(q.data, x)
}

func (q *Queue) Dequeue() int {
	if len(q.data) == 0 {
		return -1
	}
	x := q.data[0]
	q.data = q.data[1:]
	return x
}

func (q *Queue) Len() int {
	return len(q.data)
}

// Can be implemented using two queues as well
type MyStack struct {
	data Queue
}

func Constructor() MyStack {
	return MyStack{
		data: Queue{
			data: make([]int, 0),
		},
	}
}

func (this *MyStack) Push(x int) {
	this.data.Enqueue(x)
	for i := 0; i < this.data.Len()-1; i++ {
		this.data.Enqueue(this.data.Dequeue())
	}
}

func (this *MyStack) Pop() int {
	return this.data.Dequeue()
}

func (this *MyStack) Top() int {
	return this.data.data[0]
}

func (this *MyStack) Empty() bool {
	return len(this.data.data) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
