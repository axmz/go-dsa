package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{root}
	postorder := []int{}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		// almost identical to postorder
		postorder = append(postorder, node.Val)
		stack = append(stack, node.Right)
		stack = append(stack, node.Left)
	}
	return postorder
}

func preorderTraversal2(root *TreeNode) []int {
	res := []int{}

	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		} else {
			res = append(res, root.Val)
			traverse(root.Left)
			traverse(root.Right)
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
