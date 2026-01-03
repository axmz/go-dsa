package main

import (
	"fmt"
	. "godsa/utils/tree"
)

/**
 * Definition for a binary tree root.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// root.Left, root.Right = root.Right, root.Left
	root.Right = invertTree(root.Left)
	root.Left = invertTree(root.Right)

	return root
}

func main() {
	// nums := []any{4, 2, 7, 1, 3, 6, 9}
	nums := []any{1, 2}

	fmt.Println(invertTree(CreateTree(nums)))
}
