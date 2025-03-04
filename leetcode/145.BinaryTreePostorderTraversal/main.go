package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{root}
	postorder := []int{}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		postorder = append(postorder, node.Val)
		stack = append(stack, node.Left)
		stack = append(stack, node.Right)
	}
	// almost identical to preorder
	sort.SliceStable(postorder, func(i, j int) bool {
		return i > j
	})
	return postorder
}

func postorderTraversal2(root *TreeNode) []int {
	res := []int{}

	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		} else {
			traverse(root.Left)
			traverse(root.Right)
			res = append(res, root.Val)
		}
	}

	traverse(root)

	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
