package main

import "fmt"

func minimumSubarrayLength(nums []int, k int) int {
	for i := 1; i <= len(nums); i++ {
		for j := i; j <= len(nums); j++ {
			s := nums[j-i : j]
			or := 0
			for _, n := range s {
				or |= n
				if or >= k {
					return len(s)
				}
			}
		}
	}
	return -1
}

func main() {
	k := 10
	nums := []int{2, 1, 8}
	fmt.Println(minimumSubarrayLength(nums, k))
}
