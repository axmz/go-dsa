package main

import (
	"fmt"
)

type ListNode[T comparable] struct {
	Val  T
	Next *ListNode[T]
}

func reverseBetween[T comparable](head *ListNode[T], left int, right int) *ListNode[T] {
	var prev *ListNode[T] = head
	var cursor *ListNode[T] = head

	for i := 1; i <= left; {
		if i == left {
			c := right - left
			rev, after := reverseWithCount(cursor, c)
			if left == 1 {
				head = rev
			} else {
				prev.Next = rev
			}
			t := tail(head)
			t.Next = after
			break
		}

		if i < left {
			prev = cursor
			cursor = cursor.Next
			i++
		}
	}

	return head
}

func reverseWithCount[T comparable](l *ListNode[T], counter int) (rev *ListNode[T], after *ListNode[T]) {

	var next *ListNode[T] = l
	var prev *ListNode[T] = nil

	for i := 0; i <= counter; {
		cn := ListNode[T]{
			Val:  next.Val,
			Next: prev,
		}

		prev = &cn
		if i == counter {
			return prev, next.Next
		} else {
			next = next.Next
			i++
		}
	}

	return prev, next
}

func main() {

	linkedList := createLinkedListFromString("1,2,3,4,5")
	res := reverseBetween(&linkedList, 2, 4)
	s := signature(res)

	fmt.Println("Result: ", s)
}
