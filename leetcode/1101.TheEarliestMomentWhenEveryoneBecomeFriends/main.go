package main

import (
	"fmt"
	. "godsa/utils/disjointset"
	"sort"
)

func earliestAcq(logs [][]int, n int) int {
	u := NewUnion(n)
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

	connections := 0
	for _, v := range logs {
		if !u.Connected(v[1], v[2]) {
			u.Union(v[1], v[2])
			connections++
		}

		if connections == n-1 {
			return v[0]
		}
	}

	return -1
}

func main() {
	n := 0
	logs := [][]int{{20190101, 0, 1}, {20190104, 3, 4}, {20190107, 2, 3}, {20190211, 1, 5}, {20190224, 2, 4}, {20190301, 0, 3}, {20190312, 1, 2}, {20190322, 4, 5}}
	fmt.Println(earliestAcq(logs, n))
}
