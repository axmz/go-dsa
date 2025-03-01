package main

import "fmt"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	var s, s1, s2, carry int
	for l1 != nil || l2 != nil || carry != 0 {
		cur.Next = &ListNode{}
		cur = cur.Next
		if l1 == nil {
			s1 = 0
		} else {
			s1 = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			s2 = 0
		} else {
			s2 = l2.Val
			l2 = l2.Next
		}

		s = (s1 + s2 + carry) % 10
		carry = (s1 + s2 + carry) / 10

		cur.Val = s
	}

	return res.Next
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
