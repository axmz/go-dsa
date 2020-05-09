package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode = &ListNode{}
	var tail *ListNode
	if l1.Val < l2.Val {
		result.Val = l1.Val
		result.Next = l1.Next
	} else {
		result.Val = l2.Val
		result.Next = l2.Next
	}
	tail = result.Next

	for l1.Next != nil && l2.Next != nil {
		if l1.Val < l2.Val {
			tail.Val = l1.Val
			tail.Next = l1.Next
		} else {
			tail.Val = l2.Val
			tail.Next = l2.Next
		}
	}
	return result
}
