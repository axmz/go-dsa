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
	. "godsa/utils/tree"
)

// WITH POINTER
// func goodNodes(root *TreeNode) int {
// 	count := 0
// 	pathMax := -10001

// 	traverse(root, &count, pathMax)
// 	return count
// }

// func traverse(n *TreeNode, count *int, pathMax int) {

// 	if n == nil {
// 		return
// 	}

// 	if n.Val >= pathMax {
// 		pathMax = n.Val
// 		*count++
// 	}

// 	traverse(n.Left, count, pathMax)
// 	traverse(n.Right, count, pathMax)
// }

func goodNodes(root *TreeNode) int {
	count := 0
	pathMax := -10001

	return traverse(root, count, pathMax)
}

func traverse(n *TreeNode, count, pathMax int) int {

	if n == nil {
		return count
	}

	if n.Val >= pathMax {
		pathMax = n.Val
		count++
	}

	return count + (traverse(n.Left, count, pathMax) - count) + (traverse(n.Right, count, pathMax) - count)
}

func main() {

	// root := []any{4, 2, 7, 1, 3}
	// root := []any{3, 3, nil, 4, 2}
	// root := []any{1}
	root := []any{3, 1, 4, 3, nil, 1, 5}
	tree := CreateTree(root)
	res := goodNodes(tree)
	fmt.Println(res)
}
