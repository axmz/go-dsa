package main

import (
	"fmt"
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
	if k <= 0 || len(nums) == 0 || k > len(nums) {
		return false
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}

	if sum%k != 0 {
		return false
	}

	target := sum / k

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	if nums[0] > target {
		return false
	}

	buckets := make([]int, k)
	var dfs func(int) bool
	dfs = func(index int) bool {
		if index == len(nums) {
			return true
		}

		curr := nums[index]
		for i := 0; i < k; i++ {
			if buckets[i]+curr > target {
				continue
			}

			buckets[i] += curr
			if dfs(index + 1) {
				return true
			}
			buckets[i] -= curr

			// optimization: if the current bucket is empty after backtracking, no need to try other empty buckets
			if buckets[i] == 0 {
				break
			}
		}
		return false
	}

	return dfs(0)
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
