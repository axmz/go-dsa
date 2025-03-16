package main

import (
	"fmt"
	. "godsa/utils/disjointset"
	"sort"
)

func minCostToSupplyWater(n int, wells []int, pipes [][]int) int {
	costs := [][]int{}
	zeroNode := 0
	for i := 0; i < n; i++ {
		house := i + 1
		wellCost := wells[i]
		wellCostEdge := []int{zeroNode, house, wellCost}
		costs = append(costs, wellCostEdge)
	}

	costs = append(costs, pipes...)
	edges := len(costs)

	sort.Slice(costs, func(i, j int) bool {
		return costs[i][2] < costs[j][2]
	})

	u := NewUnion(edges)

	connectedEdgesCount := 0
	mstWeight := 0

	for i := 0; i < edges; i++ {
		w := costs[i][2]
		n1 := costs[i][0]
		n2 := costs[i][1]

		if !u.Connected(n1, n2) {
			u.Union(n1, n2)
			mstWeight += w
			connectedEdgesCount++
		}

		if connectedEdgesCount == edges-1 {
			break
		}
	}

	return mstWeight
}

func main() {
	n := 3
	wells := []int{1, 2, 2}
	pipes := [][]int{{1, 2, 1}, {2, 3, 1}}
	fmt.Println(minCostToSupplyWater(n, wells, pipes))
}
