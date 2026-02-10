package main

import "fmt"

func minBitwiseArray2(nums []int) []int {
	ans := make([]int, len(nums))
	for i, n := range nums {
		if n%2 == 0 { // even can't satisfy the condition: ans[i] | (ans[i] + 1) == nums[i]
			ans[i] = -1
			continue
		}
		for b := 1; n&b != 0; b <<= 1 {
			ans[i] = n - b
		}
	}
	return ans
}

func minBitwiseArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i, n := range nums {
		ans[i] = -1
		for j := 1; j < n; j++ {
			if (j | (j + 1)) == n { // ans[i] | (ans[i] + 1) == nums[i]
				ans[i] = j
				break
			}
		}
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(minBitwiseArray2(nums), nums)
}
