package main

import (
	"fmt"
	"math/bits"
	"sort"
)

func countBits2(n int) int {
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

func countBits(n int) int {
	count := 0
	for n > 0 {
		n &= n - 1
		count++
	}
	return count
}

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		bi, bj := countBits(arr[i]), countBits(arr[j])
		if bi == bj {
			return arr[i] < arr[j]
		}
		return bi < bj
	})

	return arr
}

func sortByBits2(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		bi, bj := bits.OnesCount(uint(arr[i])), bits.OnesCount(uint(arr[j]))
		if bi == bj {
			return arr[i] < arr[j]
		}
		return bi < bj
	})

	return arr
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(sortByBits(nums))
}
