package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import (
	"fmt"
	. "godsa/faang/utils"
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	if l != nil && r != nil {
		return root
	}

	if l != nil {
		return l
	}
	return r
}

func main() {
	root := []any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}
	tree := CreateTree(root)
	// p := tree.Left
	// q := tree.Right
	p := tree.Left
	q := tree.Left.Right.Right
	res := lowestCommonAncestor(tree, p, q)
	fmt.Println(res)
}
