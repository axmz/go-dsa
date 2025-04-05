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

func kthSmallest2(root *TreeNode, k int) int {
	res := root.Val
	count := k

	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {

		if root.Left != nil {
			inorder(root.Left)
		}
		count--
		if count == 0 {
			res = root.Val
			return
		}

		if root.Right != nil {
			inorder(root.Right)
		}
	}

	inorder(root)

	return res
}

func kthSmallest(root *TreeNode, k int) int {
	c := make(chan int)

	var walk func(n *TreeNode)
	walk = func(n *TreeNode) {
		if n == nil {
			return
		}
		walk(n.Left)
		c <- n.Val
		walk(n.Right)
	}

	go walk(root)

	for i := 0; i < k-1; i++ {
		<-c
	}

	return <-c
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
