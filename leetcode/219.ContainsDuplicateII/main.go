package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]bool, k)
	for i, num := range nums {
		if m[num] {
			return true
		}
		m[num] = true
		if i >= k { // clever: same as len(m) > k, basically maintaining a sliding window of size k
			delete(m, nums[i-k])
		}
	}
	return false
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
