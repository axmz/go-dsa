// the solution is simple as we only had to choose one branch where the target is
// smaller or larger than the current node value.
// note: iterative solution is cleaner
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func closestValue(root *TreeNode, target float64) int {
	return helper(root, target, root.Val)
}

func helper(node *TreeNode, target float64, closest int) int {
	if node == nil {
		return closest
	}

	diff := abs(target - float64(node.Val))
	closestDiff := abs(target - float64(closest))

	if diff < closestDiff || (diff == closestDiff && node.Val < closest) {
		closest = node.Val
	}

	if target < float64(node.Val) {
		return helper(node.Left, target, closest)
	}
	return helper(node.Right, target, closest)
}

func closestValue2(root *TreeNode, target float64) int {
	closest := root.Val
	node := root

	for node != nil {
		diff := abs(target - float64(node.Val))
		closestDiff := abs(target - float64(closest))

		if diff < closestDiff || (diff == closestDiff && node.Val < closest) {
			closest = node.Val
		}

		if target < float64(node.Val) {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	return closest
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
