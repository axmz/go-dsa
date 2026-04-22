package main

import "fmt"

func postorder(root *Node) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	for _, child := range root.Children {
		res = append(res, postorder(child)...)
	}
	res = append(res, root.Val)

	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
