/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

import (
	"fmt"
	. "godsa/utils/tree"
)

func maxLevelSum(root *TreeNode) int {
	var max, curLevel, maxLevel, sum int
	var pop *TreeNode
	max = -100001
	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) != 0 {
		curLevel++

		for i := len(q); i > 0; i-- {
			pop, q = q[0], q[1:]
			if pop.Left != nil {
				q = append(q, pop.Left)
			}
			if pop.Right != nil {
				q = append(q, pop.Right)
			}
			sum += pop.Val
		}

		if sum > max {
			max = sum
			maxLevel = curLevel
		}

		sum = 0
	}

	return maxLevel
}

func main() {
	// root := []any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}
	// root := []any{1, 7, 0, 7, -8, nil, nil}
	root := []any{-100, -200, -300, -20, -5, -10, nil}
	// root := []any{1, nil, nil}
	tree := CreateTree(root)
	res := maxLevelSum(tree)
	fmt.Println(res)
}
