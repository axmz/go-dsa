package bucket

import (
	"math"
)

func Sort(original []int) []int {
	n := len(original)
	arr := make([]int, n)
	copy(arr, original)

	min, max := arr[0], arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	// There can be various strategies to determine the number of buckets ex: n, sqrt, 24/32/64 etc.
	bucketCount := int(math.Sqrt(float64(n)))
	if bucketCount == 0 {
		bucketCount = 1
	}
	buckets := make([][]int, bucketCount)

	// Pre-allocate capacity for each bucket
	capacity := n / bucketCount
	if capacity < 1 {
		capacity = 1
	}

	for i := 0; i < bucketCount; i++ {
		buckets[i] = make([]int, 0, capacity)
	}

	// Distribute elements into buckets
	for i := 0; i < n; i++ {
		bucketIndex := (arr[i] - min) * bucketCount / (max - min + 1)
		buckets[bucketIndex] = append(buckets[bucketIndex], arr[i])
	}

	// Sort each bucket using sort algo of choice
	for i := 0; i < bucketCount; i++ {
		insertionSort(buckets[i])
	}

	res := make([]int, 0, n)
	for i := 0; i < bucketCount; i++ {
		res = append(res, buckets[i]...)
	}

	return res
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
