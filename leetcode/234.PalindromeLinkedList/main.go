package main

import (
	"fmt"
	. "godsa/utils/list"
)

func reverse(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	var temp *ListNode

	for cur != nil {
		temp = cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}

	return prev
}

func isPalindrome2(head *ListNode) bool {
	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	premiddle := slow
	middle := premiddle.Next
	newMiddle := reverse(middle)
	premiddle.Next = newMiddle

	slow = head

	for newMiddle != nil {
		if slow.Val != newMiddle.Val {
			return false
		}
		slow = slow.Next
		newMiddle = newMiddle.Next
	}

	newMiddle = reverse(premiddle.Next)
	premiddle.Next = newMiddle
	return true
}

// clever technique
func isPalindrome(head *ListNode) bool {
	frontPointer := head

	var checkRecursively func(head *ListNode) bool
	checkRecursively = func(head *ListNode) bool {
		if head != nil {
			if !checkRecursively(head.Next) {
				return false
			}
			if head.Val != (*frontPointer).Val {
				return false
			}
			(frontPointer) = frontPointer.Next
		}
		return true
	}

	return checkRecursively(head)
}

func main() {
	head := []int{1, 2, 1}
	ll := CreateLinkedList(head)
	r := isPalindrome(&ll)
	fmt.Println(r)
}
