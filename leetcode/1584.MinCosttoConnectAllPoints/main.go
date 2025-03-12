package main

import (
	"container/heap"
	"fmt"
	. "godsa/utils/disjointset"
	"sort"
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func minCostConnectPoints2(points [][]int) int {
	n := len(points)
	edges := [][]int{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			w := abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1])
			edges = append(edges, []int{w, i, j})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] < edges[j][0]
	})

	u := NewUnion(n)

	connectedEdgesCount := 0
	mstWeight := 0

	for i := 0; i < len(edges); i++ {
		w := edges[i][0]
		n1 := edges[i][1]
		n2 := edges[i][2]

		if !u.Connected(n1, n2) {
			u.Union(n1, n2)
			mstWeight += w
			connectedEdgesCount++
		}

		if connectedEdgesCount == n-1 {
			break
		}
	}

	return mstWeight
}

type Edge [3]int

type MinHeap []*Edge

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i][2] < h[j][2] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*Edge)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minCostConnectPoints(points [][]int) int {
	mstWeight := 0
	connectedNodesCount := 1
	l := len(points)

	h := &MinHeap{}
	heap.Init(h)
	for j := 1; j < l; j++ {
		w := abs(points[0][0]-points[j][0]) + abs(points[0][1]-points[j][1])
		heap.Push(h, &Edge{0, j, w})
	}

	visited := make([]bool, l)
	visited[0] = true

	for h.Len() > 0 && connectedNodesCount < l {
		top := heap.Pop(h).(*Edge)
		b := top[1]
		w := top[2]
		if visited[b] {
			continue
		}
		visited[b] = true
		mstWeight += w
		connectedNodesCount++
		for i := 0; i < l; i++ {
			if !visited[i] {
				w := abs(points[b][0]-points[i][0]) + abs(points[b][1]-points[i][1])
				heap.Push(h, &Edge{b, i, w})
			}
		}
	}

	return mstWeight
}

func main() {
	nums := [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}
	// nums := [][]int{{0, 0}, {1, 1}}
	fmt.Println(minCostConnectPoints(nums))
}
