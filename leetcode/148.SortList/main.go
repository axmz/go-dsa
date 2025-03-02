package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	var sort func(head *ListNode) *ListNode
	var merge func(headA, headB *ListNode) *ListNode

	sort = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}

		slow, fast := head, head
		for fast.Next != nil && fast.Next.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}
		listA := head
		listB := slow.Next
		slow.Next = nil

		return merge(sort(listA), sort(listB))
	}

	merge = func(headA, headB *ListNode) *ListNode {
		if headA == nil && headB == nil {
			return nil
		}

		dummy := &ListNode{}
		prev := dummy

		for headA != nil && headB != nil {
			if headA.Val <= headB.Val {
				prev.Next = headA
				headA = headA.Next
			} else {
				prev.Next = headB
				headB = headB.Next
			}
			prev = prev.Next
		}

		if headA == nil {
			prev.Next = headB
		}
		if headB == nil {
			prev.Next = headA
		}

		return dummy.Next
	}

	return sort(head)
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
