package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Node struct {
	Distance int
	Index    int
}
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance
}

func (pq *PriorityQueue) Push(a interface{}) {
	n := a.(*Node)
	*pq = append(*pq, n)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := old[len(old)-1]
	old = old[:len(old)-1]
	*pq = old
	return n
}

func adjList(graph [][]int, size int) map[int][]Node {
	list := make(map[int][]Node, size)
	for _, v := range graph {
		cur := v[0]
		dest := Node{Index: v[1], Distance: v[2]}
		list[cur] = append(list[cur], dest)
	}

	return list
}

func traverse(list map[int][]Node, time map[int]int, pq *PriorityQueue) {
	for pq.Len() > 0 {
		smallest := heap.Pop(pq).(*Node)

		if _, ok := time[smallest.Index]; ok && time[smallest.Index] <= smallest.Distance {
			continue
		}
		time[smallest.Index] = smallest.Distance
		for _, target := range list[smallest.Index] {
			heap.Push(pq, &Node{Distance: smallest.Distance + target.Distance, Index: target.Index})
		}
	}
}

func networkDelayTimeDijkstra(times [][]int, n int, k int) int {
	list := adjList(times, n)
	time := make(map[int]int)

	pq := &PriorityQueue{}
	heap.Push(pq, &Node{Distance: 0, Index: k})

	traverse(list, time, pq)

	if len(time) < n {
		return -1
	}

	m := max(time)

	return m
}

func max(m map[int]int) int {
	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}

	return max
}

func networkDelayTimeBellmanFord(times [][]int, n int, k int) int {

	distances := map[int]int{k: 0}

	for i := 0; i < n-1; i++ {
		for j := 0; j < len(times); j++ {
			src := times[j][0]

			srcDist := math.MaxInt32
			if _, ok := distances[src]; ok {
				srcDist = distances[src]
			}
			dest := times[j][1]
			dist := times[j][2]
			newDist := dist + srcDist
			if _, ok := distances[dest]; ok {
				if distances[dest] > newDist {
					distances[dest] = newDist
				}
			} else {
				distances[dest] = newDist
			}
		}
	}

	if len(distances) < n {
		return -1
	}
	return max(distances)
}

func main() {
	// t := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	// res := networkDelayTime(t, 4, 2)
	t := [][]int{{1, 2, 9}, {1, 4, 2}, {2, 5, 1}, {4, 2, 4}, {4, 5, 6}, {3, 2, 3}, {5, 3, 7}, {3, 1, 5}}
	res := networkDelayTimeBellmanFord(t, 5, 1)
	fmt.Println(res)
}
