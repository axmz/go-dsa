package main

import "fmt"

/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	n := len(grid)

	prefix := make([][]int, n+1)
	for i := range prefix {
		prefix[i] = make([]int, n+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			prefix[i+1][j+1] = prefix[i][j+1] + prefix[i+1][j] - prefix[i][j] + grid[i][j]
		}
	}

	var recurse func(x_from, x_to, y_from, y_to int) *Node
	recurse = func(x_from, x_to, y_from, y_to int) *Node {
		sum := prefix[x_to][y_to] - prefix[x_from][y_to] - prefix[x_to][y_from] + prefix[x_from][y_from]
		total := (x_to - x_from) * (y_to - y_from)
		if sum == 0 || sum == total {
			return &Node{Val: sum == total, IsLeaf: true}
		}

		x_mid := (x_from + x_to) / 2
		y_mid := (y_from + y_to) / 2

		return &Node{
			Val:         true,
			TopLeft:     recurse(x_from, x_mid, y_from, y_mid),
			TopRight:    recurse(x_from, x_mid, y_mid, y_to),
			BottomLeft:  recurse(x_mid, x_to, y_from, y_mid),
			BottomRight: recurse(x_mid, x_to, y_mid, y_to),
		}
	}

	return recurse(0, n, 0, n)
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
