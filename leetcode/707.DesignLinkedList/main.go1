package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head *Node
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

// func (this *MyLinkedList) print() {
// 	s := []int{}
// 	node := this.head
// 	for node != nil && node.next != nil {
// 		s = append(s, node.val)
// 		node = node.next
// 	}
// 	fmt.Println(s)
// }

func (this *MyLinkedList) Get(index int) int {
	// this.print()
	if index >= this.size {
		return -1
	}

	node := this.head

	if node == nil {
		return -1
	}

	for i := 0; i < index; i++ {
		node = node.next
	}

	return node.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := Node{val: val, next: this.head}
	this.head = &node
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := this.head

	if node == nil {
		this.AddAtHead(val)
		return
	}

	for node.next != nil {
		node = node.next
	}

	node.next = &Node{val: val, next: nil}
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index >= this.size {
		return
	}

	if this.head == nil {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	this.size++

	node := this.head
	var prev *Node
	for i := 0; i < index; i++ {
		prev = node
		node = node.next
	}

	prev.next = &Node{val: val, next: node}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size {
		return
	}

	this.size--

	if index == 0 {
		this.head = this.head.next
		return
	}

	node := this.head
	var prev *Node
	for i := 0; i < index; i++ {
		prev = node
		node = node.next
	}

	prev.next = node.next
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
