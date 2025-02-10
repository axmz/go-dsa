package main

import "fmt"

func sortArrayByParity(nums []int) []int {

	for i, j := 0, len(nums)-1; i < j; {
		if nums[j]%2 == 0 && nums[i]%2 == 1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
		if nums[i]%2 == 0 {
			i++
		}
		if nums[j]%2 == 1 {
			j--
		}
	}

	return nums
}

func main() {
	// nums := []int{3, 1, 2, 4}
	// nums := []int{1, 0, 3}
	nums := []int{1, 2}
	fmt.Println(sortArrayByParity(nums))
}
