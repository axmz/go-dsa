package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	prev := dummy
	curr := head

	for curr != nil && curr.Next != nil {
		next := curr.Next
		temp := next.Next

		// Swapping
		prev.Next = next
		next.Next = curr
		curr.Next = temp

		// Re-positioning pointers
		prev = curr
		curr = temp
	}

	return dummy.Next
}

func main() {

	fmt.Println(swapPairs(nil))
}
