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
func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}

	var i, j *ListNode = head, head.Next
	for {
		if j.Next == nil || j.Next.Next == nil {
			break
		} else {
			j = j.Next.Next
			i = i.Next
		}
	}

	if i.Next.Next == nil {
		i.Next = nil
	} else {
		i.Next = i.Next.Next
	}

	return head
}

func main() {
	// head := []int{1, 3, 4, 7, 1, 2, 6}
	head := []int{1}

	ll := CreateLinkedList(head)
	newll := deleteMiddle(&ll)
	fmt.Println(newll.Signature())
}
