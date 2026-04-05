package main

import "fmt"

func mergeTriplets(triplets [][]int, target []int) bool {
	good := make(map[int]bool, 0)
	for _, triplet := range triplets {
		if triplet[0] > target[0] || triplet[1] > target[1] || triplet[2] > target[2] {
			continue
		}

		if triplet[0] == target[0] {
			good[0] = true
		}

		if triplet[1] == target[1] {
			good[1] = true
		}

		if triplet[2] == target[2] {
			good[2] = true
		}

		if len(good) == 3 {
			return true
		}
	}

	return false
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
