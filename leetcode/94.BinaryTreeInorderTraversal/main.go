package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal2(root *TreeNode) []int {
	res := []int{}

	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		} else {
			traverse(root.Left)
			res = append(res, root.Val)
			traverse(root.Right)
		}
	}

	traverse(root)

	return res
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	curr := root
	for curr != nil || len(stack) != 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curr.Val)
		curr = curr.Right
	}
	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
