package main

import "fmt"

func insert(aNode *Node, x int) *Node {
	newNode := &Node{
		Val: x,
	}

	if aNode == nil {
		newNode.Next = newNode
		return newNode
	}

	for ptr := aNode.Next; true; ptr = ptr.Next {
		next := ptr.Next // clever, instead of ptr = next; next = next.next you have one line here and ptr = ptr.next in the for loop line

		if aNode == ptr ||
			(ptr.Val <= x && x <= next.Val) ||
			(next.Val < ptr.Val && (x > ptr.Val || x < next.Val)) {
			newNode.Next = next
			ptr.Next = newNode
			break
		}
	}

	return aNode
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
