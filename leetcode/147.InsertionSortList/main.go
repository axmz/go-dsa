package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{0, nil}
	cur := head

	for cur != nil {
		prev := dummy
		for prev.Next != nil && cur.Val > prev.Next.Val {
			prev = prev.Next
		}
		next := cur.Next
		temp := prev.Next
		prev.Next = cur
		cur.Next = temp
		cur = next
	}

	return dummy.Next
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
