package main

import (
	"container/heap"
	"fmt"
	"math"
)

func findCheapestPrice2(n int, flights [][]int, src int, dst int, k int) int {
	adj := make(map[int][][]int, n)
	for _, v := range flights {
		s := v[0]
		d := v[1]
		w := v[2]
		adj[s] = append(adj[s], []int{d, w})
	}

	var pop []int
	Q1 := [][]int{{src, 0}}
	Q2 := make([][]int, 0)
	lvl := 0

	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt
	}

	for lvl <= k && len(Q1) > 0 {
		pop, Q1 = Q1[0], Q1[1:]
		s := pop[0]
		w := pop[1]
		for _, v := range adj[s] {
			d := v[0]
			nw := v[1] + w
			if dist[d] < nw {
				continue
			}
			dist[d] = nw
			Q2 = append(Q2, []int{d, nw})
		}

		if len(Q1) == 0 {
			lvl++
			Q1 = Q2
		}
	}

	if dist[dst] == math.MaxInt {
		return -1
	}
	return dist[dst]
}

type Node struct {
	destination int
	weight      int
	level       int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
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

func findCheapestPrice3(n int, flights [][]int, src int, dst int, k int) int {
	adj := make(map[int][][]int, n)
	for _, v := range flights {
		s := v[0]
		d := v[1]
		w := v[2]
		adj[s] = append(adj[s], []int{d, w})
	}

	stops := make([]int, n)
	for i := range stops {
		stops[i] = math.MaxInt
	}

	pq := &PriorityQueue{}
	heap.Push(pq, &Node{src, 0, 0})

	for pq.Len() > 0 {
		top := heap.Pop(pq).(*Node)
		d := top.destination
		w := top.weight
		lvl := top.level

		if lvl > k+1 || stops[d] < lvl {
			continue
		}

		if d == dst {
			return w
		}

		stops[d] = lvl

		for _, v := range adj[d] {
			dn := v[0]
			wn := v[1]
			heap.Push(pq, &Node{dn, wn + w, lvl + 1})
		}
	}

	if stops[dst] == math.MaxInt {
		return -1
	}

	return stops[dst]
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt
	}

	dist[src] = 0

	for i := 0; i <= k; i++ {
		temp := make([]int, n)
		copy(temp, dist)
		for _, f := range flights {
			s := f[0]
			d := f[1]
			c := f[2]
			if dist[s] == math.MaxInt {
				continue
			}

			newCost := dist[s] + c
			if newCost < dist[d] {
				temp[d] = newCost
			}
		}

		copy(dist, temp)
	}

	if dist[dst] == math.MaxInt {
		return -1
	}

	return dist[dst]
}

func main() {
	n := 15
	flights := [][]int{{10, 14, 43}, {1, 12, 62}, {4, 2, 62}, {14, 10, 49}, {9, 5, 29}, {13, 7, 53}, {4, 12, 90}, {14, 9, 38}, {11, 2, 64}, {2, 13, 92}, {11, 5, 42}, {10, 1, 89}, {14, 0, 32}, {9, 4, 81}, {3, 6, 97}, {7, 13, 35}, {11, 9, 63}, {5, 7, 82}, {13, 6, 57}, {4, 5, 100}, {2, 9, 34}, {11, 13, 1}, {14, 8, 1}, {12, 10, 42}, {2, 4, 41}, {0, 6, 55}, {5, 12, 1}, {13, 3, 67}, {3, 13, 36}, {3, 12, 73}, {7, 5, 72}, {5, 6, 100}, {7, 6, 52}, {4, 7, 43}, {6, 3, 67}, {3, 1, 66}, {8, 12, 30}, {8, 3, 42}, {9, 3, 57}, {12, 6, 31}, {2, 7, 10}, {14, 4, 91}, {2, 3, 29}, {8, 9, 29}, {2, 11, 65}, {3, 8, 49}, {6, 14, 22}, {4, 6, 38}, {13, 0, 78}, {1, 10, 97}, {8, 14, 40}, {7, 9, 3}, {14, 6, 4}, {4, 8, 75}, {1, 6, 56}}
	src := 1
	dst := 4
	k := 10
	fmt.Println(findCheapestPrice(n, flights, src, dst, k))

	n = 4
	flights = [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}
	src = 0
	dst = 3
	k = 1
	fmt.Println(findCheapestPrice(n, flights, src, dst, k))

	n = 5
	flights = [][]int{{4, 1, 1}, {1, 2, 3}, {0, 3, 2}, {0, 4, 10}, {3, 1, 1}, {1, 4, 3}}
	src = 2
	dst = 1
	k = 1
	fmt.Println(findCheapestPrice(n, flights, src, dst, k))

	n = 4
	flights = [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}}
	src = 0
	dst = 3
	k = 1
	fmt.Println(findCheapestPrice(n, flights, src, dst, k))
}
