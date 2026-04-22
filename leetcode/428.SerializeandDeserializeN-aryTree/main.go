package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Codec struct {
}

func Constructor() *Codec {
	return &Codec{}
}

func (this *Codec) serialize(root *Node) string {
	if root == nil {
		return ""
	}

	tokens := make([]string, 0)
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}

		tokens = append(tokens, strconv.Itoa(node.Val), strconv.Itoa(len(node.Children)))
		// tokens look like: [val, childCount, val, childCount, ...]
		// [1, 3, 2, 0, 3, 0, 4, 0] represents a node with value 1 and 3 children (2, 3, 4), each with 0 children
		for _, child := range node.Children {
			dfs(child)
		}
		// There are other approaches:
		// - values [1,2,3,4,5,6], parents [-1,0,0,0,2,2]
		// - 1 | 2 3 4 | # | 5 6 | # # #
		// - 1(2,3(5,6),4)
	}

	dfs(root)
	return strings.Join(tokens, " ")
}

func (this *Codec) deserialize(data string) *Node {
	if data == "" {
		return nil
	}

	tokens := strings.Fields(data)
	idx := 0

	var build func() *Node
	build = func() *Node {
		if idx >= len(tokens) {
			return nil
		}

		val, _ := strconv.Atoi(tokens[idx])
		idx++
		childCount, _ := strconv.Atoi(tokens[idx])
		idx++

		node := &Node{Val: val, Children: make([]*Node, 0, childCount)}
		for i := 0; i < childCount; i++ {
			node.Children = append(node.Children, build())
		}

		return node
	}

	return build()
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
