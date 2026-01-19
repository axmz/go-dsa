package main

import (
	"fmt"
	"sort"
)

func smallestDistancePair(numbers []int, k int) int {
	sort.Ints(numbers)
	minDistance, maxDistance := 0, numbers[len(numbers)-1]-numbers[0]

	for minDistance < maxDistance {
		midDistance := minDistance + (maxDistance-minDistance)/2
		if countPairsWithinDistance(numbers, midDistance) < k {
			minDistance = midDistance + 1
		} else {
			maxDistance = midDistance
		}
	}

	return minDistance
}

func countPairsWithinDistance(numbers []int, targetDistance int) int {
	count, left := 0, 0
	for right := 1; right < len(numbers); right++ {
		for numbers[right]-numbers[left] > targetDistance {
			left++
		}
		count += right - left
	}
	return count
}

func main() {
	x := 3
	nums := []int{1, 2, 4, 7, 12, 13}
	fmt.Println(smallestDistancePair(nums, x))
}
