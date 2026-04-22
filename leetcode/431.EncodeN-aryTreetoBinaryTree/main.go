package main

import "fmt"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() *Codec {
	return &Codec{}
}

// This is LCRS (Left Child Right Sibling) representation of an N-ary tree in a binary tree.
// The left child of a binary tree node represents the first child of the N-ary tree node,
// and the right child of a binary tree node represents the next sibling of the N-ary tree node.
func (this *Codec) encode(root *Node) *TreeNode {
	if root == nil {
		return nil
	}

	treeNode := &TreeNode{Val: root.Val}
	if len(root.Children) > 0 {
		treeNode.Left = this.encode(root.Children[0])
		current := treeNode.Left
		for i := 1; i < len(root.Children); i++ {
			current.Right = this.encode(root.Children[i])
			current = current.Right
		}
	}

	return treeNode
}

func (this *Codec) decode(root *TreeNode) *Node {
	if root == nil {
		return nil
	}

	node := &Node{Val: root.Val}
	current := root.Left
	for current != nil {
		child := this.decode(current)
		node.Children = append(node.Children, child)
		current = current.Right
	}

	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * bst := obj.encode(root);
 * ans := obj.decode(bst);
 */

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
