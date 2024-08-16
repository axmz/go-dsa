package main

import (
	"fmt"
	. "godsa/leetcode/utils/list"
)

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	var odd, oddHead, prev *ListNode
	cur := head
	isEven := false

	for cur != nil {
		if isEven {
			if oddHead == nil {
				oddHead = cur
				odd = oddHead
			} else {
				odd.Next = cur
				odd = odd.Next
			}
			if cur.Next != nil {
				prev.Next = cur.Next
			}
		} else {
			prev = cur
		}

		isEven = !isEven
		cur = cur.Next
	}

	odd.Next = nil
	prev.Next = oddHead

	return head
}

func main() {
	// head := []int{1, 2, 3, 4, 5}
	head := []int{1, 2, 3, 4, 5}
	ll := CreateLinkedList(head)
	r := oddEvenList(&ll)
	fmt.Println(r.Signature(), "xxxx")
}
