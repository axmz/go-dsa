package main

import "fmt"

type ListNode[T comparable] struct {
	Val  T
	Next *ListNode[T]
}

func hasCycleSimpleAlgo(head *ListNode[string]) bool {
	set := make(map[*ListNode[string]]bool)
	for {
		if head == nil {
			return false
		}
		if head.Next == nil {
			return false
		}
		if _, ok := set[head]; ok {
			return true
		} else {
			set[head] = true
			head = head.Next
		}
	}
}

func detectCycleFloydAlgo(head *ListNode[string]) *ListNode[string] {
	turtle := head
	hare := head

	if head == nil {
		return nil
	}

	for {

		if hare.Next == nil || hare.Next.Next == nil {
			return nil
		}

		hare = hare.Next.Next
		turtle = turtle.Next

		if *hare == *turtle {
			turtle = head

			for *hare != *turtle {
				hare = hare.Next
				turtle = turtle.Next
			}

			return hare
		}

	}

}

func main() {

	linkedList := createLinkedListFromString("3,2,0,-4")
	// linkedList := ListNode[string]{}
	t := tail[string](&linkedList)
	t.Next = linkedList.Next

	s := hasCycleSimpleAlgo(&linkedList)
	s2 := detectCycleFloydAlgo(&linkedList)

	fmt.Println("Result: ", s, s2)
}
