package main

import (
	"fmt"
	. "godsa/utils/list"
)

func oddEvenList(head *ListNode) *ListNode {
	var cur *ListNode = head
	var prev *ListNode = nil
	oddHead := &ListNode{}
	oddTail := oddHead
	i := 0

	if cur == nil {
		return nil
	}

	for cur != nil {
		if i%2 == 0 {
			prev = cur
		} else {
			prev.Next = cur.Next
			cur.Next = nil
			oddTail.Next = cur
			oddTail = oddTail.Next
		}
		cur = prev.Next
		i++
	}

	prev.Next = oddHead.Next
	return head
}

func oddEvenList2(head *ListNode) *ListNode {
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
