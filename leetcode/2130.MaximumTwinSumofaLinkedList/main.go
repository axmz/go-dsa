package main

import (
	"fmt"
	. "godsa/utils/list"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
	// MIDDLE
	var fast, slow *ListNode
	fast, slow = head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// REVERSE
	var rev, cur, next *ListNode

	cur = slow // 4
	for cur != nil {
		next = cur.Next // 5,6,7
		cur.Next = nil  // (4)
		cur.Next = rev  // 4,nil
		rev = cur       // 4,nil
		cur = next      // 5,6,7
	}

	// MAX
	var sum, max int
	cur = rev
	for cur != nil {
		sum = head.Val.(int) + cur.Val.(int)
		if sum > max {
			max = sum
		}
		cur = cur.Next
		head = head.Next
	}

	return max
}

func main() {
	// head := []int{4, 2, 2, 3}
	// head := []int{5, 4, 2, 1}
	head := []int{1, 2, 3, 4, 5, 6}
	ll := CreateLinkedList(head)
	r := pairSum(&ll)
	fmt.Println(">>>", r)
}
