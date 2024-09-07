package main

import (
	"fmt"
	. "godsa/utils/list"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var revHead *ListNode
	for head != nil {
		head.Next, revHead, head = revHead, head, head.Next
	}
	return revHead
}

func recurse(head *ListNode) (*ListNode, *ListNode) {
	if head.Next == nil {
		return head, head
	}

	h, lastNode := recurse(head.Next)
	lastNode.Next = head
	return h, lastNode.Next
}

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	h, lastNode := recurse(head)
	lastNode.Next = nil
	return h
}

func reverseListIterative(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	// No need for cur, you can just use head instead
	var temp, prev, cur *ListNode
	cur = head

	for cur.Next != nil {
		temp = cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}

	// return prev is enough
	cur.Next = prev
	return cur
}

func main() {
	l := CreateIntLinkedListFromString("1,2,3,4,5")

	r := reverseList(&l)
	s := r.Signature()
	fmt.Println(s)
}

// 1 > 2 > 3 > 4 > 5
// 5 > 4 > 3 > 2 > 1
