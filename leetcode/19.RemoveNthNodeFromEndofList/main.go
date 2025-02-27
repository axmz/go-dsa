package main

import (
	"fmt"
	. "godsa/utils/list"
)

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	node := head
	l := 1
	for node.Next != nil {
		node = node.Next
		l++
	}

	if l == n {
		return head.Next
	}

	node = head
	for i := 0; i < l-n; i++ {
		node = node.Next
	}
	if node.Next != nil {
		node.Next = node.Next.Next
	}

	return node
}

func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}

	node := dummy
	l := 0
	for node.Next != nil {
		node = node.Next
		l++
	}

	node = dummy
	for i := 0; i < l-n; i++ {
		node = node.Next
	}
	node.Next = node.Next.Next

	return dummy.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}

	first := dummy
	second := dummy
	for i := n; i > 0; i-- {
		first = first.Next
	}

	for first.Next != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next

	return dummy.Next
}

func main() {
	s := "1,2,3,4,5"
	n := 2
	head := CreateIntLinkedListFromString(s)
	fmt.Println(removeNthFromEnd(&head, n))
}
