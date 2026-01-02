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

func generate(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	// if ok, v := memo;
	res := []*TreeNode{}
	for i := start; i <= end; i++ {
		leftTrees := generate(start, i-1)
		rightTrees := generate(i+1, end)
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				root := &TreeNode{Val: i}
				root.Left = left
				root.Right = right
				res = append(res, root)
			}
		}
	}

	return res
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	return generate(1, n)
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
