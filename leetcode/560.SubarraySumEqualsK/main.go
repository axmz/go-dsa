package main

import "fmt"

func subarraySum(nums []int, k int) int {
	cumulativeSum := make([]int, len(nums))
	sum, count := 0, 0
	m := make(map[int]int, len(nums))
	m[0] = 1

	for i, v := range nums {

		sum += v
		cumulativeSum[i] = sum

		dif := sum - k

		if v, ok := m[dif]; ok {
			count += v
		}

		m[sum]++
	}

	return count
}

func main() {

	// nums := []int{1, 1, 1}
	// k := 2
	// nums := []int{3, 4, 7, 2, -3, 1, 4, 2}
	// k := 7
	nums := []int{1}
	k := 0
	// nums := []int{1, 2, 3}
	// k := 3

	res := subarraySum(nums, k)
	fmt.Println(res)
}
