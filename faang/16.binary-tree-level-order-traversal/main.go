package main

import (
	"fmt"
	. "godsa/utils/tree"
)

func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	Q1 := []*TreeNode{root}
	Q2 := []*TreeNode{}

	for i := 0; len(Q1) != 0; i++ {
		if i == len(Q1) {
			res = append(res, flat(Q1))
			Q1 = Q2
			Q2 = []*TreeNode{}
			i = -1
		} else {
			l := Q1[i].Left
			r := Q1[i].Right
			if l != nil {
				Q2 = append(Q2, l)
			}

			if r != nil {
				Q2 = append(Q2, r)
			}
		}
	}

	return res
}

func flat(node []*TreeNode) (res []int) {
	for _, v := range node {
		res = append(res, v.Val)
	}

	return
}

func levelOrder3(root *TreeNode) [][]int {
	res := [][]int{}

	var helper func(root *TreeNode, level int)
	helper = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if len(res) == level {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)
		helper(root.Left, level+1)
		helper(root.Right, level+1)
	}

	helper(root, 0)

	return res
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)
	var pop *TreeNode

	for len(q) > 0 {
		size := len(q)
		level := []int{}

		for i := 0; i < size; i++ {
			pop, q = q[0], q[1:]
			level = append(level, pop.Val)
			if pop.Left != nil {
				q = append(q, pop.Left)
			}
			if pop.Right != nil {
				q = append(q, pop.Right)
			}
		}

		res = append(res, level)
	}

	return res
}

func main() {
	root := []interface{}{3, 9, 20, nil, nil, 15, 7}
	tree := CreateTree(root)
	res := levelOrder(tree)
	fmt.Println(res)
}
