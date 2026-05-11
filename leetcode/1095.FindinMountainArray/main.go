package main

import "fmt"

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
	l := mountainArr.length()

	// 1. Find peak index
	lo, hi := 0, l-1
	for lo < hi {
		mid := (hi-lo)/2 + lo
		if mountainArr.get(mid) < mountainArr.get(mid+1) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	peak := lo

	// 2. Binary search ascending left side [0, peak]
	lo, hi = 0, peak
	for lo <= hi {
		mid := (hi-lo)/2 + lo
		v := mountainArr.get(mid)
		if v == target {
			return mid
		} else if v < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	// 3. Binary search descending right side [peak+1, l-1]
	lo, hi = peak+1, l-1
	for lo <= hi {
		mid := (hi-lo)/2 + lo
		v := mountainArr.get(mid)
		if v == target {
			return mid
		} else if v > target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return -1
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
