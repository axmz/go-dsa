package main

import (
	"fmt"
)

// horrible, very slow.
type pair struct {
	i   int
	num string
}

func radixSort(arr []string, k int, trim int) pair {
	pairs := make([]pair, len(arr))
	for i, num := range arr {
		pairs[i] = pair{i, num}
	}

	for i := 0; i < trim; i++ {
		count := [10]int{}
		for _, p := range pairs {
			digit := p.num[len(p.num)-i-1] - '0'
			count[digit]++
		}

		newArr := make([]pair, len(arr))
		for j := 1; j < 10; j++ {
			count[j] += count[j-1]
		}

		for j := len(pairs) - 1; j >= 0; j-- {
			num := pairs[j].num
			digit := num[len(num)-i-1] - '0'
			count[digit]--
			newArr[count[digit]] = pairs[j]
		}
		pairs = newArr
	}

	return pairs[k-1]
}

func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
	res := make([]int, 0, len(queries))

	for _, q := range queries {
		k, trim := q[0], q[1]
		sorted := radixSort(nums, k, trim)
		res = append(res, sorted.i)
	}

	return res
}

func main() {
	nums := []string{"102", "473", "251", "814"}
	queries := [][]int{{1, 1}, {2, 3}, {4, 2}, {1, 2}}
	fmt.Println(smallestTrimmedNumbers(nums, queries))
}
