package main

import "fmt"

func quickSelect(nums []int, k int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		pivotIndex := partition(nums, left, right)

		if pivotIndex == k {
			return nums[pivotIndex]
		} else if pivotIndex < k {
			left = pivotIndex + 1
		} else {
			right = pivotIndex - 1
		}
	}

	return -1
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

func wiggleSort(nums []int) {
	n := len(nums)

	fmt.Println(nums)
	// Step 1: find median and partition around it
	median := quickSelect(nums, n/2)

	fmt.Println(nums, median)
	// Virtual index mapping
	mapIndex := func(i int) int {
		return (1 + 2*i) % (n | 1)
	}

	// Step 2: 3-way partition
	left, i, right := 0, 0, n-1

	for i <= right {
		if nums[mapIndex(i)] > median {
			nums[mapIndex(left)], nums[mapIndex(i)] =
				nums[mapIndex(i)], nums[mapIndex(left)]
			left++
			i++
		} else if nums[mapIndex(i)] < median {
			nums[mapIndex(right)], nums[mapIndex(i)] =
				nums[mapIndex(i)], nums[mapIndex(right)]
			right--
		} else {
			i++
		}
	}
}

func main() {
	nums := []int{5, 1, 1, 7, 6}
	// nums := []int{1, 2, 3, 4, 5}
	wiggleSort(nums)
}
