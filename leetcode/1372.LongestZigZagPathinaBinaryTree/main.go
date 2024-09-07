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

const (
	Left  = -1
	Right = 1
)

func recurse(root *TreeNode, c int, prevDir int) int {
	if root == nil {
		return c
	}

	var cl, cr int
	c++
	nextDir := -1 * prevDir

	if nextDir == Left {
		cl = c
		cr = 0

	} else if nextDir == Right {
		cr = c
		cl = 0
	}

	return max(recurse(root.Left, cl, Left), recurse(root.Right, cr, Right))
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func longestZigZag(root *TreeNode) int {

	c := 0
	if root == nil {
		return c
	}

	return max(recurse(root.Left, c, Left), recurse(root.Right, c, Right))
}

func main() {
	// root := []any{3, 1, 4, 3, nil, 1, 5}
	root := []any{1, 1, 1, nil, 1, nil, nil, 1, 1, nil, 1}
	// root := []any{1, nil, 1, 1, 1, nil, nil, 1, 1, nil, 1, nil, nil, nil, 1}
	tree := CreateTree(root)
	res := longestZigZag(tree)
	fmt.Println(res)
}
