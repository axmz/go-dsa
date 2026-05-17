package main

import (
	"fmt"
	. "godsa/utils/disjointset"
	"sort"
)

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	edgesWithIndex := make([][]int, len(edges))
	for i, e := range edges {
		edgesWithIndex[i] = []int{e[0], e[1], e[2], i}
	}

	// sort by weight
	sort.Slice(edgesWithIndex, func(i, j int) bool {
		return edgesWithIndex[i][2] < edgesWithIndex[j][2]
	})

	u := NewUnion(n)
	mst := 0
	for _, e := range edgesWithIndex {
		if !u.Connected(e[0], e[1]) {
			u.Union(e[0], e[1])
			mst += e[2]
		}
	}

	critical := make([]int, 0)
	pseudoCritical := make([]int, 0)
	for _, e := range edgesWithIndex {
		u := NewUnion(n)
		cost := 0
		edgesUsed := 0
		// Critical edge test
		// Critical means without the edge being tested, the MST cost is more or graph is disconnected
		for _, other := range edgesWithIndex {
			// skip the edge being tested
			if other[3] == e[3] {
				continue
			}

			// build MST without the edge being tested
			if !u.Connected(other[0], other[1]) {
				u.Union(other[0], other[1])
				cost += other[2]
				edgesUsed++
			}
		}

		// if graph is disconnected or cost > mst, then the edge being tested is critical
		if edgesUsed != n-1 || cost > mst {
			critical = append(critical, e[3])
			continue
		}

		// Pseudo-Critical edge test
		// Pseudo-Critical means with or without the edge being tested, the MST cost is the same
		u = NewUnion(n)
		u.Union(e[0], e[1])
		cost = e[2]
		edgesUsed = 1
		for _, other := range edgesWithIndex {
			if other[3] == e[3] {
				continue
			}

			if !u.Connected(other[0], other[1]) {
				u.Union(other[0], other[1])
				cost += other[2]
				edgesUsed++
			}
		}

		// if graph is connected and cost == mst, then the edge being tested is pseudo-critical
		if edgesUsed == n-1 && cost == mst {
			pseudoCritical = append(pseudoCritical, e[3])
		}
	}

	return [][]int{critical, pseudoCritical}
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
