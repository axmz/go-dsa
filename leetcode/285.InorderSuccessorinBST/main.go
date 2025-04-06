package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor2(root *TreeNode, p *TreeNode) *TreeNode {
	stack := []*TreeNode{root}
	flag := false

	for {
		for root != nil {
			if root.Left != nil {
				stack = append(stack, root.Left)
			}
			root = root.Left
		}
		if len(stack) == 0 {
			return nil
		}
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if flag {
			return pop
		}

		if pop == p {
			flag = true
		}

		if pop.Right != nil {
			stack = append(stack, pop.Right)
		}
		root = pop.Right
	}
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var succ *TreeNode

	for root != nil {
		if p.Val < root.Val {
			succ = root
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return succ
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
