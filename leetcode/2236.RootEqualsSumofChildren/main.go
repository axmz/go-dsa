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
func checkTree(root *TreeNode) bool {
	if root.Val == (root.Left.Val + root.Right.Val) {
		return true
	}
	return false
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
