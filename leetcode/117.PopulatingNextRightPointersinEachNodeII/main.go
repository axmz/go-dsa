package main

import (
	"fmt"
	. "godsa/utils/tree"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect2(root *Node) *Node {
	if root == nil {
		return root
	}
	q1, q2 := []*Node{root}, []*Node{}
	var pop *Node

	for len(q1) > 0 {
		pop, q1 = q1[0], q1[1:]

		if len(q1) > 0 {
			pop.Next = q1[0]
		}

		if pop.Left != nil {
			q2 = append(q2, pop.Left)
		}

		if pop.Right != nil {
			q2 = append(q2, pop.Right)
		}

		if len(q1) == 0 {
			q1, q2 = q2, q1
		}
	}

	return root
}

func connect(root *Node) *Node {
	leftmost := root
	head := leftmost
	parent := &Node{}

	for leftmost != nil {
		if parent == nil {
			parent = head
			leftmost = leftmost.Left
			head = leftmost
		} else if parent.Right != head {
			head.Next = parent.Right
			head = head.Next
		} else if parent.Next != nil {
			head.Next = parent.Next.Left
			head = head.Next
			parent = parent.Next
		} else {
			parent = leftmost
			leftmost = leftmost.Left
			head = leftmost
		}
	}

	return root
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	curr := root
	for curr != nil {
		dummy := &Node{} // Dummy for next level’s head
		head := dummy
		// Traverse current level
		for parent := curr; parent != nil; parent = parent.Next {
			if parent.Left != nil {
				head.Next = parent.Left
				head = head.Next
			}
			if parent.Right != nil {
				head.Next = parent.Right
				head = head.Next
			}
		}
		curr = dummy.Next // Move to next level
	}
	return root
}

func main() {
	root := []any{1, 2, 3, 4, 5, 6, 7, 8}
	res := connect(CreateTree(root))
	fmt.Println(res)
}
