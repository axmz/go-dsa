package main

import "fmt"

func reverseList(head *ListNode) *ListNode {
	node := head
	var prev *ListNode = nil
	for node != nil {
		temp := node.Next
		node.Next = prev
		prev = node
		node = temp
	}
	return prev
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	reversedHalf := reverseList(slow.Next)
	slow.Next = nil

	curr := head
	for curr != nil && reversedHalf != nil {
		nextCurr := curr.Next
		nextReversed := reversedHalf.Next
		curr.Next = reversedHalf
		reversedHalf.Next = nextCurr
		curr = nextCurr
		reversedHalf = nextReversed
	}
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
