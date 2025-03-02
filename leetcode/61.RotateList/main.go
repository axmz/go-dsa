package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	l, r := head, head
	len := 1

	for i := 0; i < k && r.Next != nil; i++ {
		r = r.Next
		len++
	}

	if r.Next == nil {
		k = k % len
		if k == 0 {
			return head
		}

		r.Next = head
		for i := 0; i <= k; i++ {
			r = r.Next
		}

		for i := len - k; i > 0; i-- {
			l = l.Next
			r = r.Next
		}
	} else {
		for r.Next != nil {
			l = l.Next
			r = r.Next
		}
		r.Next = head
	}

	head = l.Next
	l.Next = nil

	return head
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
