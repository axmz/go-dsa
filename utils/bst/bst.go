package bst

import (
	"encoding/json"
	"fmt"
)

// BST structure
type BST struct {
	Root *Node
}

// Insert inserts the root node or passes the insert responsibility to the next node
func (t *BST) Insert(num int) {
	if t.Root == nil {
		t.Root = &Node{num, nil, nil}
	} else {
		t.Root.Insert(num)
	}
}

// Lookup passes the search to the root node
func (t BST) Lookup(num int) (node *Node, parentNode *Node, err error) {
	if t.Root == nil {
		return node, parentNode, fmt.Errorf("Has no root")
	}
	return t.Root.Lookup(num, t.Root)
}

// Remove removes the node from the tree
func (t *BST) Remove(num int) (node *Node, parentNode *Node, err error) {
	node, parentNode, err = t.Lookup(num)

	if err != nil {
		println(err.Error())
		return node, parentNode, err
	}

	if node == parentNode { // if root case
		t.Root = nil
	} else if node.Right == nil && node.Left == nil { // no right, no left - remove entirely
		if parentNode.Right.Value == node.Value {
			parentNode.Right = nil
		}
		if parentNode.Left.Value == node.Value {
			parentNode.Left = nil
		}
	} else if node.Right == nil { // no right - connect parent with left
		if parentNode.Right.Value == node.Value {
			parentNode.Right = node.Left
		}
		if parentNode.Left.Value == node.Value {
			parentNode.Left = node.Left
		}
	} else if node.Right != nil { // right - connect parent with leftmost
		leftmost := node.Right
		for {
			if leftmost.Left != nil {
				leftmost = leftmost.Left
			} else {
				break
			}
		}
		if parentNode.Right.Value == node.Value {
			parentNode.Right = leftmost
		}
		if parentNode.Left.Value == node.Value {
			parentNode.Left = leftmost
		}
	}

	return node, parentNode, err
}

// Print - prints the entire tree in json
func (t BST) Print() {
	printTree, _ := json.MarshalIndent(t, "", "|  ")
	fmt.Println(string(printTree))
}

// NewBst initializes an empty Bst
func NewBst() *BST {
	return &BST{
		Root: nil,
	}
}

// BFS traverses the tree
func (t *BST) BFS() []int {
	queue := []*Node{t.Root}
	list := []int{}
	for len(queue) > 0 {
		f := queue[0]
		if f.Left != nil {
			queue = append(queue, f.Left)
		}
		if f.Right != nil {
			queue = append(queue, f.Right)
		}
		list = append(list, f.Value)
		queue = queue[1:]
	}
	return list
}

// BFSR recursively traverses the tree
func (t *BST) BFSR(queue []*Node, list []int) []int {
	if queue == nil {
		queue = []*Node{t.Root}
		list = []int{}
	}

	if len(queue) == 0 {
		return list
	}

	f := queue[0]
	if f.Left != nil {
		queue = append(queue, f.Left)
	}
	if f.Right != nil {
		queue = append(queue, f.Right)
	}
	list = append(list, f.Value)
	queue = queue[1:]

	return t.BFSR(queue, list)
}

// InOrder dfs in order traversal
func (t *BST) InOrder(n *Node, list *[]int) *[]int {
	// if ( n == nil) {
	// 	return list
	// }
	println("case0")

	if n.Left != nil {
		println("case1")
		t.InOrder(n.Left, list)
	}

	println("case3")
	*list = append(*list, n.Value)
	fmt.Println(list)

	if n.Right != nil {
		t.InOrder(n.Right, list)
	}

	return list
}

// PreOrder dfs pre order traversal
func (t *BST) PreOrder(n *Node, list *[]int) *[]int {
	*list = append(*list, n.Value)

	if n.Left != nil {
		t.PreOrder(n.Left, list)
	}

	if n.Right != nil {
		t.PreOrder(n.Right, list)
	}

	return list
}

// PostOrder dfs pre order traversal
func (t *BST) PostOrder(n *Node, list *[]int) *[]int {
	if n.Left != nil {
		t.PostOrder(n.Left, list)
	}

	if n.Right != nil {
		t.PostOrder(n.Right, list)
	}

	*list = append(*list, n.Value)

	return list
}
