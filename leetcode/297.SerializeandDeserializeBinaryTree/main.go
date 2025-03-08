package main

import (
	"fmt"
	. "godsa/utils/tree"
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {

	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	nodes := []any{}

	var traverse func(*TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			nodes = append(nodes, nil)
			return
		}

		q := []*TreeNode{root}
		var pop *TreeNode
		for len(q) > 0 {
			pop, q = q[0], q[1:]
			if pop != nil {
				nodes = append(nodes, pop.Val)
				q = append(q, pop.Left, pop.Right)
			} else {
				nodes = append(nodes, nil)
			}
		}

	}

	removeNils := func() {
		lastDigit := -1
		for i, v := range nodes {
			if v != nil {
				lastDigit = i
			}
		}

		if lastDigit > 0 {
			nodes = nodes[:lastDigit+1]
		}
	}

	traverse(root)
	removeNils()
	return fmt.Sprintf("%v", nodes)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	root := &TreeNode{}
	if data == "" {
		return nil
	}
	nodes := []*TreeNode{}
	removeBrackets := func() {
		data = data[1 : len(data)-1]
	}

	convertAToNodes := func() {
		strnodes := strings.Split(data, " ")
		for _, v := range strnodes {
			num, err := strconv.Atoi(v)
			if err != nil {
				nodes = append(nodes, nil)
			} else {
				nodes = append(nodes, &TreeNode{Val: num})
			}
		}
	}

	buildTree := func() {
		var q []*TreeNode
		if nodes[0] == nil {
			return
		}
		pop := nodes[0]
		nodes := nodes[1:]
		root.Val = pop.Val
		q = []*TreeNode{root}

		for len(q) > 0 {
			pop, q = q[0], q[1:]
			if pop == nil {
				continue
			}
			if len(nodes) > 0 {
				pop.Left = nodes[0]
				q = append(q, nodes[0])
				nodes = nodes[1:]
			}
			if len(nodes) > 0 {
				pop.Right = nodes[0]
				q = append(q, nodes[0])
				nodes = nodes[1:]
			}
		}
	}

	removeBrackets()
	convertAToNodes()
	buildTree()

	return root
}

func main() {
	ser := Constructor()
	deser := Constructor()
	root := []any{1, 2, 3, 4, 5, 6, 7, 8}
	// root := []any{1, 2, 3, nil, nil, 4, 5}
	data := ser.serialize(CreateTree(root))
	ans := deser.deserialize(data)
	fmt.Println(data)
	fmt.Println(ans)
}
