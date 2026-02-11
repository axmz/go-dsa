package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	cur, next := head, head.Next
	for next != nil {
		if cur.Val == next.Val {
			cur.Next = next.Next
			next = cur.Next
		} else {
			cur = next
			next = next.Next
		}
	}
	return head
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
