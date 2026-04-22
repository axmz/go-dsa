package main

import "fmt"

func preorder(root *Node) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	res = append(res, root.Val)
	for _, child := range root.Children {
		res = append(res, preorder(child)...)
	}

	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
