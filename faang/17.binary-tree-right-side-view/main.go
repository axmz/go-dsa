package main

import (
	"fmt"
	. "godsa/faang/utils"
)

func rightSideViewDFS(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	level := 0

	traverse(root, level, &res)

	return res
}

func traverse(node *TreeNode, level int, res *[]int) {
	if level > len(*res)-1 {
		*res = append(*res, node.Val)
	}

	if node.Right != nil {
		traverse(node.Right, level+1, res)
	}
	if node.Left != nil {
		traverse(node.Left, level+1, res)
	}
}

func rightSideViewBFS(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	Q1 := []*TreeNode{root}
	Q2 := []*TreeNode{}

	for i := 0; len(Q1) != 0; i++ {
		if i == len(Q1) {
			res = append(res, rightmost(Q1))
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

func rightmost(node []*TreeNode) (res int) {
	return node[len(node)-1].Val
}

func main() {
	root := []interface{}{1, 2, 3, nil, 5, nil, 4}
	// root := []interface{}{3, 9, 20, nil, nil, 15, 7}
	tree := CreateTree(root)
	res := rightSideViewDFS(tree)
	res2 := rightSideViewBFS(tree)
	fmt.Println(res, res2)
}
