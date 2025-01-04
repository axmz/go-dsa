package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func lengthOfLIS2(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	curMax := 1

	for i := len(nums) - 2; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
				curMax = max(curMax, dp[i])
			}
		}
	}

	return curMax
}

func binarySearch(a []int, x int) int {
	start, mid, end := 0, 0, len(a)-1
	for start != end {
		mid = (start + end) >> 1
		switch {
		case a[mid] >= x:
			end = mid
		case a[mid] < x:
			start = mid + 1
		}
	}
	return start
}

func lengthOfLIS(nums []int) int {
	lsi := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] > lsi[len(lsi)-1] {
			lsi = append(lsi, nums[i])
		} else {
			idx := binarySearch(lsi, nums[i])
			if idx != -1 {
				lsi[idx] = nums[i]
			}
		}
	}

	return len(lsi)
}

// 1,2,4,6,8
func main() {
	// nums := []int{0, 1, 0, 3, 2, 3}
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}
