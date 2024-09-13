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

// Level Order traversal
// func connect(root *Node) *Node {
// 	if root == nil {
// 		return root
// 	}

// 	var pop *Node
// 	queue := []*Node{root}
// 	for len(queue) > 0 {
// 		q := []*Node{}
// 		prev := &Node{}
// 		for len(queue) > 0 {
// 			pop, queue = queue[0], queue[1:]
// 			if prev != nil {
// 				prev.Next = pop
// 			}
// 			prev = pop
// 			if pop != nil {
// 				if pop.Left != nil {
// 					q = append(q, pop.Left)
// 				}
// 				if pop.Right != nil {
// 					q = append(q, pop.Right)
// 				}
// 			}
// 		}
// 		prev.Next = nil
// 		queue = append(queue, q...)
// 	}

// 	return root
// }

// O(1) pointers traversal
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

func main() {
	root := []any{1, 2, 3, 4, 5, 6, 7, 8}
	res := connect(CreateTree(root))
	fmt.Println(res)
}
