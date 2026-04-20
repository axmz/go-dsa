package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value int
	row   int
	col   int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].value < pq[j].value }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.row = priority
	heap.Fix(pq, item.index)
}

// priority queue approach
func kthSmallest2(matrix [][]int, k int) int {
	l := len(matrix)
	pq := make(PriorityQueue, l)
	for i, items := range matrix {
		pq[i] = &Item{
			value: items[0],
			row:   i,
			col:   0,
		}
	}
	heap.Init(&pq)

	for i := 0; i < k-1; i++ {
		item := heap.Pop(&pq).(*Item)
		if item.col < l-1 {
			heap.Push(&pq, &Item{
				value: matrix[item.row][item.col+1],
				row:   item.row,
				col:   item.col + 1,
			})
		}
	}

	return heap.Pop(&pq).(*Item).value
}

// binary search on value approach
func kthSmallest(matrix [][]int, k int) int {
	l := len(matrix)
	left, right := matrix[0][0], matrix[l-1][l-1]

	for left < right {
		mid := left + (right-left)/2
		count := 0
		// here we count how many numbers are less than mid
		for r := 0; r < l; r++ {
			// if last number in the row is less than mid
			for c := l - 1; c >= 0; c-- {
				if matrix[r][c] <= mid {
					// then we add the entire row to count
					count += c + 1
					break
				}
			}
		}

		// after inner loop we check how many numbers we got so far < k
		// if not enough then we expand our search to the right half
		if count < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
