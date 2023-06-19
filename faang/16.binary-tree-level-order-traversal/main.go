package main

import (
	"fmt"
	. "godsa/faang/utils"
)

func levelOrder(root *TreeNode) [][]int {
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

func main() {
	root := []interface{}{3, 9, 20, nil, nil, 15, 7}
	tree := CreateTree(root)
	res := levelOrder(tree)
	fmt.Println(res)
}
