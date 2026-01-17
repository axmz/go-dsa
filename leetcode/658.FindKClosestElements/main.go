// nice clean solutions:
// https://leetcode.com/problems/find-k-closest-elements/solutions/6056237/easy-and-straightforward-beats-100-o1-me-qtmw/

package main

import (
	"fmt"
	"sort"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Using native sort with custom comparator
func findClosestElements2(arr []int, k int, x int) []int {
	sort.Slice(arr, func(i, j int) bool {
		if abs(arr[i]-x) == abs(arr[j]-x) { // why we need this check?
			return arr[i] < arr[j]
		}
		return abs(arr[i]-x) < abs(arr[j]-x) // why this is not enough?
	})
	result := arr[:k]
	sort.Ints(result)
	return result
}

func findClosestElements(arr []int, k int, x int) []int {
	// Binary search to find position closest to x
	left, right := 0, len(arr)-1
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] < x {
			left = mid + 1 // left will end up the idx we are looking for
		} else {
			right = mid
		}
	}

	// Two pointers expanding from closest position
	i, j := left-1, left
	for count := 0; count < k; count++ {
		if i < 0 {
			j++
		} else if j >= len(arr) || x-arr[i] <= arr[j]-x {
			i--
		} else {
			j++
		}
	}

	return arr[i+1 : j]
}

// finding left lower bound
func findClosestElements1(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-k
	for left < right {
		mid := left + (right-left)/2
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return arr[left : left+k]
}

func main() {
	x := -30
	k := 4
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(findClosestElements(nums, k, x))
}
