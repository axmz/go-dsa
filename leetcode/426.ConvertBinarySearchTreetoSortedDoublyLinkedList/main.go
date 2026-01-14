package main

import "fmt"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}

	dummy := &Node{}
	tail := dummy

	var recurse func(node *Node)
	recurse = func(node *Node) {
		if node.Left != nil {
			recurse(node.Left)
		}

		tail.Right = node
		node.Left = tail
		tail = node

		if node.Right != nil {
			recurse(node.Right)
		}
	}

	recurse(root)
	tail.Right = dummy.Right
	dummy.Right.Left = tail

	return dummy.Right
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
