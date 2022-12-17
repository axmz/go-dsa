package main

import (
	"fmt"
	"strings"
)

func reverseLinkedList[T comparable](l ListNode[T]) ListNode[T] {

	var next ListNode[T] = l
	var prev *ListNode[T] = nil

	for {
		cn := ListNode[T]{
			Val:  next.Val,
			Next: prev,
		}

		prev = &cn
		if next.Next == nil {
			break
		} else {
			next = *next.Next
		}
	}

	return *prev
}

func traverse[T comparable](l ListNode[T]) {
	fmt.Println("\n-------------------")
	for {
		fmt.Printf("\n TRAVERSE %v", l.Val)
		if l.Next == nil {
			break
		}
		l = *l.Next
	}
}

func compare[T comparable](a *ListNode[T], b *ListNode[T]) bool {
	for a.Next != nil || b.Next != nil {
		if a.Val != b.Val {
			return false
		} else {
			a = a.Next
			b = b.Next
		}
	}
	return a.Next == b.Next
}

func signature[T comparable](l *ListNode[T]) string {
	var sb strings.Builder

	for {
		sb.WriteString(fmt.Sprint(l.Val, ","))
		if l.Next == nil {
			break
		}
		l = l.Next
	}

	return strings.TrimSuffix(sb.String(), ",")
}

func createLinkedListFromString(s string) ListNode[string] {
	strs := strings.Split(s, ",")
	return createLinkedList(strs)
}

func createLinkedList[T comparable](l []T) ListNode[T] {
	var linkedList ListNode[T]
	var lastNode *ListNode[T] = &linkedList

	for i, v := range l {
		newNode := ListNode[T]{
			Val:  v,
			Next: nil,
		}

		if i == 0 {
			linkedList = newNode
		} else {
			lastNode.Next = &newNode
			lastNode = &newNode
		}
	}

	return linkedList
}

func tail[T comparable](l *ListNode[T]) *ListNode[T] {
	cn := l
	for cn.Next != nil {
		cn = cn.Next
	}
	return cn
}
