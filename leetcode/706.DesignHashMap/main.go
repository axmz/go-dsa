package main

import "fmt"

type MyHashMap struct {
	buckets []*Node
	size    int
}

func Constructor() MyHashMap {
	return MyHashMap{buckets: make([]*Node, 769), size: 769}
}

func (this *MyHashMap) _hash(key int) int {
	return key % this.size
}

func (this *MyHashMap) Put(key int, value int) {
	h := this._hash(key)
	if this.buckets[h] == nil {
		this.buckets[h] = NewNode(key, value)
		return
	}
	this.buckets[h].Insert(key, value)
}

func (this *MyHashMap) Get(key int) int {
	h := this._hash(key)
	if this.buckets[h] == nil {
		return -1
	}
	if kv, _, err := this.buckets[h].Lookup(key, nil); err != nil {
		return -1
	} else {
		return kv.Value
	}

}

func (this *MyHashMap) Remove(key int) {
	h := this._hash(key)
	if this.buckets[h] != nil {
		this.buckets[h] = this.buckets[h].Remove(key)
	}

}

// Node structure of the node
type Node struct {
	Key   int
	Value int
	Left  *Node
	Right *Node
}

func NewNode(key, value int) *Node {
	return &Node{Key: key, Value: value}
}

// Insert inserts a value into BST
func (n *Node) Insert(key, value int) {
	if key == n.Key {
		n.Value = value
		return
	}
	if key > n.Key {
		if n.Right == nil {
			n.Right = &Node{Key: key, Value: value}
		} else {
			n.Right.Insert(key, value)
		}
	} else {
		if n.Left == nil {
			n.Left = &Node{Key: key, Value: value}
		} else {
			n.Left.Insert(key, value)
		}
	}
}

// Lookup searches for a value in BST and returns the node and its parent
func (n *Node) Lookup(key int, parent *Node) (*Node, *Node, error) {
	if n == nil {
		return nil, parent, fmt.Errorf("node not found")
	}
	if n.Key == key {
		return n, parent, nil
	}
	if key > n.Key {
		return n.Right.Lookup(key, n)
	}
	return n.Left.Lookup(key, n)
}

// Remove removes the node with the given value from the tree and returns the new root
func (n *Node) Remove(key int) *Node {
	if n == nil {
		return nil
	}

	if key < n.Key {
		n.Left = n.Left.Remove(key)
		return n
	} else if key > n.Key {
		n.Right = n.Right.Remove(key)
		return n
	}

	// Found the node to remove
	// Case 1: No children
	if n.Left == nil && n.Right == nil {
		return nil
	}
	// Case 2: Only right child
	if n.Left == nil {
		return n.Right
	}
	// Case 3: Only left child
	if n.Right == nil {
		return n.Left
	}

	// Case 4: Two children - find inorder successor (smallest in right subtree)
	minRight := n.Right
	for minRight.Left != nil {
		minRight = minRight.Left
	}
	// Replace current node's key and value with successor's key and value
	n.Key = minRight.Key
	n.Value = minRight.Value
	// Remove the successor node
	n.Right = n.Right.Remove(minRight.Key)
	return n
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
