package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	current := head
	for current != nil && current.Next != nil {
		next := current.Next
		gcdNode := &ListNode{Val: gcd(current.Val, next.Val)}
		current.Next = gcdNode
		gcdNode.Next = next
		current = next
	}
	return head
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
