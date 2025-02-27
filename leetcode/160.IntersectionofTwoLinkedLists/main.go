package main

import "fmt"

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	A := headA
	B := headB

	for A != B {
		if A.Next != nil {
			A = A.Next
		} else {
			A = headB
		}

		if B.Next != nil {
			B = B.Next
		} else {
			B = headA
		}
	}

	return A
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
