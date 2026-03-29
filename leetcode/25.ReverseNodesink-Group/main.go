package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func checkKNodes(head *ListNode, k int) bool {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return false
		}
		node = node.Next
	}
	return true
}

func reverseList(head *ListNode, k int) (*ListNode, *ListNode, *ListNode) {
	node := head
	end := head
	var prev *ListNode = nil
	for i := 0; i < k; i++ {
		temp := node.Next
		node.Next = prev
		prev = node
		node = temp
	}
	return prev, end, node
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 1 {
		return head
	}

	node := head
	dummy := &ListNode{Next: nil}
	prevTail := dummy
	for checkKNodes(node, k) {
		reversed, end, next := reverseList(node, k)
		prevTail.Next = reversed
		prevTail = end
		node = next
	}
	prevTail.Next = node
	return dummy.Next
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
