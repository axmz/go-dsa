package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList2(head *Node) *Node {
	var recurse func(*Node) *Node
	visited := make(map[*Node]*Node)
	recurse = func(head *Node) *Node {
		if head == nil {
			return head
		}

		if _, ok := visited[head]; ok {
			return visited[head]
		}

		n := &Node{Val: head.Val}
		visited[head] = n
		n.Next = recurse(head.Next)
		n.Random = recurse(head.Random)
		return n
	}

	return recurse(head)
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		n := &Node{Val: cur.Val, Next: cur.Next}
		cur.Next = n
		cur = n.Next
	}

	cur = head
	for cur != nil {
		cr := cur.Random
		if cr != nil {
			cur.Next.Random = cr.Next
		}
		cur = cur.Next.Next
	}

	oldHead := head
	newHead := head.Next
	cur = oldHead
	curNew := newHead

	for cur != nil {
		cur.Next = cur.Next.Next
		cur = cur.Next
		if curNew.Next != nil {
			curNew.Next = curNew.Next.Next
			curNew = curNew.Next
		}
	}

	return newHead
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
