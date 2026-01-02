package bst

import (
	"encoding/json"
	"fmt"
)

// TreeNode structure of the node
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// Insert inserts a value into BST
func (n *TreeNode) Insert(num int) {
	if n == nil {
		return
	}
	if num > n.Value {
		if n.Right == nil {
			n.Right = &TreeNode{Value: num}
		} else {
			n.Right.Insert(num)
		}
	} else {
		if n.Left == nil {
			n.Left = &TreeNode{Value: num}
		} else {
			n.Left.Insert(num)
		}
	}
}

// Lookup searches for a value in BST and returns the node and its parent
func (n *TreeNode) Lookup(num int, parent *TreeNode) (*TreeNode, *TreeNode, error) {
	if n == nil {
		return nil, parent, fmt.Errorf("node not found")
	}
	if n.Value == num {
		return n, parent, nil
	}
	if num > n.Value {
		return n.Right.Lookup(num, n)
	}
	return n.Left.Lookup(num, n)
}

// Remove removes the node with the given value from the tree
func (n *TreeNode) Remove(num int) error {
	node, parent, err := n.Lookup(num, nil)
	if err != nil {
		return err
	}

	// Can't remove root without returning new root
	if parent == nil {
		return fmt.Errorf("cannot remove root node directly")
	}

	var replacement *TreeNode

	// Determine replacement node
	if node.Left == nil && node.Right == nil {
		replacement = nil // Leaf node
	} else if node.Right == nil {
		replacement = node.Left // Only left child
	} else if node.Left == nil {
		replacement = node.Right // Only right child
	} else {
		// Both children exist: find inorder successor (leftmost in right subtree)
		replacement = node.Right
		for replacement.Left != nil {
			replacement = replacement.Left
		}
	}

	// Update parent's pointer
	if parent.Left == node {
		parent.Left = replacement
	} else {
		parent.Right = replacement
	}

	return nil
}

// Print - prints the node in json
func (n *TreeNode) Print() {
	printTree, _ := json.MarshalIndent(n, "", "|  ")
	fmt.Println(string(printTree))
}

// BFS traverses the tree breadth-first
func (n *TreeNode) BFS() []int {
	if n == nil {
		return []int{}
	}
	queue := []*TreeNode{n}
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

// BFSR recursively traverses the tree breadth-first
func (n *TreeNode) BFSR() []int {
	if n == nil {
		return []int{}
	}
	return n.bfsrHelper([]*TreeNode{n}, []int{})
}

func (n *TreeNode) bfsrHelper(queue []*TreeNode, list []int) []int {
	if len(queue) == 0 {
		return list
	}

	first := queue[0]
	queue = queue[1:]

	if first.Left != nil {
		queue = append(queue, first.Left)
	}
	if first.Right != nil {
		queue = append(queue, first.Right)
	}
	list = append(list, first.Value)

	return n.bfsrHelper(queue, list)
}

// InOrder dfs in order traversal (Left, Root, Right)
func (n *TreeNode) InOrder() []int {
	if n == nil {
		return []int{}
	}
	result := []int{}
	result = append(result, n.Left.InOrder()...)
	result = append(result, n.Value)
	result = append(result, n.Right.InOrder()...)
	return result
}

// PreOrder dfs pre order traversal (Root, Left, Right)
func (n *TreeNode) PreOrder() []int {
	if n == nil {
		return []int{}
	}
	result := []int{}
	result = append(result, n.Value)
	result = append(result, n.Left.PreOrder()...)
	result = append(result, n.Right.PreOrder()...)
	return result
}

// PostOrder dfs post order traversal (Left, Right, Root)
func (n *TreeNode) PostOrder() []int {
	if n == nil {
		return []int{}
	}
	result := []int{}
	result = append(result, n.Left.PostOrder()...)
	result = append(result, n.Right.PostOrder()...)
	result = append(result, n.Value)
	return result
}

// NewTreeNode creates a new TreeNode with the given value
func NewTreeNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

// Example usage (commented out - uncomment and change package to main to run)
/*
func main() {
	// Create root node
	tree := NewTreeNode(0)
	tree.Insert(9)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(20)
	tree.Insert(170)
	tree.Insert(15)
	tree.Insert(1)

	// Example: Remove node
	// err := tree.Remove(1)
	// if err != nil {
	//     fmt.Println("Error:", err)
	// }

	// Example: Lookup node
	// node, parent, err := tree.Lookup(4, nil)
	// if err == nil {
	//     fmt.Println("Found node:", node.Value)
	//     if parent != nil {
	//         fmt.Println("Parent:", parent.Value)
	//     }
	// }

	// tree.Print()

	// BFS
	// bfs := tree.BFS()
	// fmt.Println("BFS:", bfs)

	// BFSR
	// bfsr := tree.BFSR()
	// fmt.Println("BFSR:", bfsr)

	// DFS
	inOrder := tree.InOrder()
	preOrder := tree.PreOrder()
	postOrder := tree.PostOrder()
	fmt.Println("InOrder:", inOrder)
	fmt.Println("PreOrder:", preOrder)
	fmt.Println("PostOrder:", postOrder)
}
*/
