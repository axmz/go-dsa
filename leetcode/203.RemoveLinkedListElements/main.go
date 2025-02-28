package main

import (
	"fmt"
	. "godsa/utils/list"
)

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	cur := dummy.Next

	for cur != nil {
		if cur.Val == val {
			prev.Next = cur.Next
			cur = cur.Next
		} else {
			prev = cur
			cur = cur.Next
		}
	}

	return dummy.Next
}

func main() {
	head := []int{1, 2, 6, 3, 4, 5, 6}
	val := 6
	ll := CreateLinkedList(head)
	r := removeElements(&ll, val)
	fmt.Println(r.Signature(), "xxxx")
}
