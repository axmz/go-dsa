package main

import (
	"fmt"
)

func pivotV2(nums []int) int {
	i := 0
	high := len(nums) - 1
	j := high
	pivot := nums[0]

	for i < j {
		for i < high && nums[i] <= pivot {
			i++
		}
		for j > 0 && nums[j] > pivot {
			j--
		}
		if i < j {
			nums[j], nums[i] = nums[i], nums[j]
		}
	}
	nums[j], nums[0] = nums[0], nums[j]
	return j
}

func quickSortV2(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	p := pivotV2(nums)
	quickSortV2(nums[:p])
	quickSortV2(nums[p+1:])
	return nums
}

func pivotV1(nums []int) int {
	l := len(nums)
	p := l - 1
	pivotElement := nums[p]
	i := 0

	for j := 0; j < p; j++ {
		if nums[j] < pivotElement {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[p] = nums[p], nums[i]
	return i
}

func quickSortV1(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	p := pivotV1(nums)
	quickSortV1(nums[:p])
	quickSortV1(nums[p+1:])
	return nums
}

// V1 and V2 are not ideal solutions
// b := []int{3, 2, 1, 5, 6, 4}
func pivotV3(nums []int, low, high int) int {
	pivot := nums[high]
	i := low

	for j := low; j < high; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[high] = nums[high], nums[i]
	return i
}

func quickSortV3(nums []int, low, high int) []int {
	if high-low < 1 {
		return nums
	}
	p := pivotV3(nums, low, high)
	quickSortV3(nums, low, p-1)
	quickSortV3(nums, p+1, high)
	return nums
}

func hoares(nums []int, k, low, high int) int {
	p := pivotV3(nums, low, high)
	t := len(nums) - k
	if p == t {
		return nums[t]
	}
	if p < t {
		return hoares(nums, k, p+1, high)
	} else {
		return hoares(nums, k, 0, p-1)
	}
}

func findKthLargest(nums []int, k int) int {
	n := hoares(nums, k, 0, len(nums)-1)
	return n
}

func main() {
	a := []int{8, 9, 7, 2, 6, 7, 3}
	b := []int{3, 2, 1, 5, 6, 4}
	r := findKthLargest(b, 2)
	r2 := quickSortV3(a, 0, len(a)-1)
	fmt.Println(r, r2)
}
