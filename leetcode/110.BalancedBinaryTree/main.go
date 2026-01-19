package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + max(height(node.Left), height(node.Right))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftHeight := height(root.Left)
	rightHeight := height(root.Right)

	if abs(leftHeight-rightHeight) > 1 {
		return false
	}

	// This is necessary because the heights can be equal but the subtrees unbalanced
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
