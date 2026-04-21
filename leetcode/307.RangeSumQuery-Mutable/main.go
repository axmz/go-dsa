package main

import (
	"fmt"
)

// // sqrt decomposition
// type NumArray struct {
// 	nums      []int
// 	blockSums []int
// 	blockSize int
// }

// func Constructor(nums []int) NumArray {
// 	n := len(nums)
// 	s := math.Sqrt(float64(n))
// 	blockSize := int(s)
// 	blocksCount := n/blockSize + 1
// 	blockSums := make([]int, blocksCount)

// 	for i := 0; i < n; i++ {
// 		blockSums[i/blockSize] += nums[i]
// 	}

// 	return NumArray{
// 		nums:      nums,
// 		blockSums: blockSums,
// 		blockSize: blockSize,
// 	}
// }

// func (this *NumArray) Update(index int, val int) {
// 	blockIndex := index / this.blockSize
// 	this.blockSums[blockIndex] += val - this.nums[index]
// 	this.nums[index] = val
// }

// func (this *NumArray) SumRange(left int, right int) int {
// 	sum := 0
// 	for i := left; i <= right; {
// 		if i%this.blockSize != 0 || i+this.blockSize-1 > right {
// 			sum += this.nums[i]
// 			i++
// 		} else {
// 			sum += this.blockSums[i/this.blockSize]
// 			i += this.blockSize
// 		}
// 	}
// 	return sum
// }

// // segment tree
// type Node struct {
// 	start, end int
// 	sum        int
// 	left       *Node
// 	right      *Node
// }

// func buildSegmentTree(nums []int, start, end int) *Node {
// 	if start > end {
// 		return nil
// 	}
// 	if start == end {
// 		return &Node{start: start, end: end, sum: nums[start]}
// 	}
// 	mid := start + (end-start)/2
// 	left := buildSegmentTree(nums, start, mid)
// 	right := buildSegmentTree(nums, mid+1, end)
// 	return &Node{
// 		start: start,
// 		end:   end,
// 		sum:   left.sum + right.sum,
// 		left:  left,
// 		right: right,
// 	}
// }

// type NumArray struct {
// 	root *Node
// }

// func Constructor(nums []int) NumArray {
// 	return NumArray{
// 		root: buildSegmentTree(nums, 0, len(nums)-1),
// 	}
// }

// func (this *NumArray) Update(index int, val int) {
// 	var update func(node *Node)
// 	update = func(node *Node) {
// 		if node.start == node.end {
// 			node.sum = val
// 			return
// 		}
// 		mid := node.start + (node.end-node.start)/2
// 		if index <= mid {
// 			update(node.left)
// 		} else {
// 			update(node.right)
// 		}
// 		node.sum = node.left.sum + node.right.sum
// 	}
// 	update(this.root)
// }

// func (this *NumArray) SumRange(left int, right int) int {
// 	var sumRange func(node *Node, left, right int) int
// 	sumRange = func(node *Node, left, right int) int {
// 		if node.start == left && node.end == right {
// 			return node.sum
// 		}
// 		mid := node.start + (node.end-node.start)/2
// 		if right <= mid {
// 			return sumRange(node.left, left, right)
// 		} else if left > mid {
// 			return sumRange(node.right, left, right)
// 		} else {
// 			return sumRange(node.left, left, mid) + sumRange(node.right, mid+1, right)
// 		}
// 	}
// 	return sumRange(this.root, left, right)
// }

// // segment tree - array based
// type NumArray struct {
// 	tree []int
// 	n    int
// }

// func Constructor(nums []int) NumArray {
// 	n := len(nums)
// 	segmentTree := NumArray{
// 		tree: make([]int, 4*n),
// 		n:    n,
// 	}
// 	if n == 0 {
// 		return segmentTree
// 	}

// 	var build func(treeIndex, left, right int)
// 	build = func(treeIndex, left, right int) {
// 		if left == right {
// 			segmentTree.tree[treeIndex] = nums[left]
// 			return
// 		}

// 		mid := left + (right-left)/2
// 		leftChild := treeIndex * 2
// 		rightChild := leftChild + 1

// 		build(leftChild, left, mid)
// 		build(rightChild, mid+1, right)
// 		segmentTree.tree[treeIndex] = segmentTree.tree[leftChild] + segmentTree.tree[rightChild]
// 	}

// 	build(1, 0, n-1)
// 	return segmentTree
// }

// func (this *NumArray) Update(index int, val int) {
// 	if this.n == 0 {
// 		return
// 	}

// 	var update func(treeIndex, left, right int)
// 	update = func(treeIndex, left, right int) {
// 		if left == right {
// 			this.tree[treeIndex] = val
// 			return
// 		}

// 		mid := left + (right-left)/2
// 		leftChild := treeIndex * 2
// 		rightChild := leftChild + 1

// 		if index <= mid {
// 			update(leftChild, left, mid)
// 		} else {
// 			update(rightChild, mid+1, right)
// 		}
// 		this.tree[treeIndex] = this.tree[leftChild] + this.tree[rightChild]
// 	}

// 	update(1, 0, this.n-1)
// }

// func (this *NumArray) SumRange(left int, right int) int {
// 	if this.n == 0 {
// 		return 0
// 	}

// 	var sumRange func(treeIndex, segmentLeft, segmentRight, queryLeft, queryRight int) int
// 	sumRange = func(treeIndex, segmentLeft, segmentRight, queryLeft, queryRight int) int {
// 		if segmentLeft == queryLeft && segmentRight == queryRight {
// 			return this.tree[treeIndex]
// 		}

// 		mid := segmentLeft + (segmentRight-segmentLeft)/2
// 		leftChild := treeIndex * 2
// 		rightChild := leftChild + 1

// 		if queryRight <= mid {
// 			return sumRange(leftChild, segmentLeft, mid, queryLeft, queryRight)
// 		}
// 		if queryLeft > mid {
// 			return sumRange(rightChild, mid+1, segmentRight, queryLeft, queryRight)
// 		}

// 		return sumRange(leftChild, segmentLeft, mid, queryLeft, mid) +
// 			sumRange(rightChild, mid+1, segmentRight, mid+1, queryRight)
// 	}

// 	return sumRange(1, 0, this.n-1, left, right)
// }

// binary indexed tree - fenwick tree
type NumArray struct {
	nums []int
	sums []int
}

func Constructor(nums []int) NumArray {
	size := len(nums)
	sums := make([]int, size)
	for i, num := range nums {
		sums[i] += num
		if j := i | (i + 1); j < size {
			sums[j] += sums[i]
		}
	}
	return NumArray{
		nums: nums,
		sums: sums,
	}
}

func (this *NumArray) Update(index int, val int) {
	delta := val - this.nums[index]
	this.nums[index] = val
	for size := len(this.sums); index < size; index |= index + 1 {
		this.sums[index] += delta
	}
}
func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.Sum(right)
	}
	return this.Sum(right) - this.Sum(left-1)
}
func (this *NumArray) Sum(right int) int {
	sum := 0
	for right >= 0 {
		sum += this.sums[right]
		right = (right & (right + 1)) - 1
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
