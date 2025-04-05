package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		size := len(q)
		arr := []int{}

		for i := 0; i < size; i++ {
			pop := q[0]
			q = q[1:]

			flip := len(res)%2 == 0

			if flip {
				arr = append(arr, pop.Val)
			} else {
				arr = append([]int{pop.Val}, arr...)
			}

			if pop.Left != nil {
				q = append(q, pop.Left)
			}

			if pop.Right != nil {
				q = append(q, pop.Right)
			}
		}
		res = append(res, arr)
	}

	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
