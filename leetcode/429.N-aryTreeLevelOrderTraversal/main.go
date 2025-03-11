package main

import "fmt"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	res := [][]int{}

	if root == nil {
		return res
	}

	var pop *Node
	Q1 := []*Node{root}
	Q2 := []*Node{}

	res = append(res, []int{root.Val})

	for len(Q1) > 0 {
		pop, Q1 = Q1[0], Q1[1:]

		Q2 = append(Q2, pop.Children...)

		if len(Q1) == 0 {
			values := []int{}
			for _, v := range Q2 {
				values = append(values, v.Val)
			}
			if len(values) > 0 {
				res = append(res, values)
			}
			Q1, Q2 = Q2, Q1
		}
	}

	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
