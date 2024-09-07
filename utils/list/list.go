package list

import (
	"fmt"
	"strconv"
	"strings"
)

type Listable interface {
	Signature() string
}

type ListNode struct {
	Val  any
	Next *ListNode
}

func (l *ListNode) Signature() string {
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

func CreateIntLinkedListFromString(s string) ListNode {
	strs := strings.Split(s, ",")
	ints := make([]int, len(strs))

	for i, s := range strs {
		ints[i], _ = strconv.Atoi(s)
	}
	return CreateLinkedList(ints)
}

func CreateLinkedListFromString(s string) ListNode {
	strs := strings.Split(s, ",")
	return CreateLinkedList(strs)
}

func CreateLinkedList[T comparable](l []T) ListNode {
	var linkedList ListNode
	var lastNode *ListNode = &linkedList

	for i, v := range l {
		newNode := ListNode{
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
