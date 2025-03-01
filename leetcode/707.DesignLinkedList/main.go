package main

import "fmt"

type Node struct {
	val  int
	next *Node
	prev *Node
}

type MyLinkedList struct {
	sentinel *Node // Dummy node linking head and tail
	size     int
}

func Constructor() MyLinkedList {
	sentinel := &Node{}      // Sentinel node
	sentinel.next = sentinel // Points to itself when empty
	sentinel.prev = sentinel
	return MyLinkedList{sentinel: sentinel}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.size {
		return -1
	}
	node := this.sentinel.next // Start at head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node.val
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	this.size++
	curr := this.sentinel
	for i := 0; i < index; i++ { // Stop at prev node
		curr = curr.next
	}
	newNode := &Node{val: val, next: curr.next, prev: curr}
	curr.next.prev = newNode
	curr.next = newNode
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size {
		return
	}
	this.size--
	curr := this.sentinel
	for i := 0; i < index; i++ { // Stop at prev node
		curr = curr.next
	}
	curr.next = curr.next.next // Skip target node
	curr.next.prev = curr
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
