package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func (this *BSTIterator) traverseLeft(node *TreeNode) {
	cur := node
	for cur != nil {
		this.stack = append(this.stack, cur)
		cur = cur.Left
	}
}

func Constructor(root *TreeNode) BSTIterator {
	stack := []*TreeNode{}
	cur := root
	for cur != nil {
		stack = append(stack, cur)
		cur = cur.Left
	}
	return BSTIterator{
		stack,
	}
}

func (this *BSTIterator) Next() int {
	n := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.traverseLeft(n.Right)
	return n.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
