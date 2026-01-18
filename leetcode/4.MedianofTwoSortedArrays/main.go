package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// X : [....minLeftX | maxRightX....]
	// Y : [....minLeftY | maxRightY....]
	m, n := len(nums1), len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}
	l, r := 0, m
	for l <= r {
		maxLeftX := math.MinInt64
		minRightX := math.MaxInt64
		maxLeftY := math.MinInt64
		minRightY := math.MaxInt64

		partitionX := l + (r-l)/2
		partitionY := (m+n+1)/2 - partitionX // remaining elements

		if partitionX > 0 {
			maxLeftX = nums1[partitionX-1]
		}
		if partitionX < m {
			minRightX = nums1[partitionX]
		}
		if partitionY > 0 {
			maxLeftY = nums2[partitionY-1]
		}
		if partitionY < n {
			minRightY = nums2[partitionY]
		}

		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			if (m+n)%2 == 0 {
				return float64(max(maxLeftX, maxLeftY)+min(minRightX, minRightY)) / 2.0
			} else {
				return float64(max(maxLeftX, maxLeftY))
			}
		} else if maxLeftX > minRightY {
			r = partitionX - 1
		} else {
			l = partitionX + 1
		}
	}

	return 0.0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums1 := []int{1, 2, 3, 4, 6, 8, 10}
	nums2 := []int{5, 7, 9, 11, 12}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
