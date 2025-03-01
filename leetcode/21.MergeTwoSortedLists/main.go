package main

import "fmt"

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	prev := dummy

	for {
		if list1 == nil && list2 == nil {
			break
		} else if list1 == nil {
			prev.Next = list2
			break
		} else if list2 == nil {
			prev.Next = list1
			break
		} else if list1.Val < list2.Val {
			temp := list1.Next
			prev.Next = list1
			list1 = temp
		} else {
			temp := list2.Next
			prev.Next = list2
			list2 = temp
		}
		prev = prev.Next
	}

	return dummy.Next
}

func main() {
	fmt.Println(x, nums)
}
