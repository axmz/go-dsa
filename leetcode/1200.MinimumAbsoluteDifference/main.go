package main

import "fmt"

func minimumAbsDifference(arr []int) [][]int {
	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	rangeSize := (max - min) + 1
	r := make([]bool, rangeSize)
	offset := -min

	for _, v := range arr {
		r[v+offset] = true
	}
	result := [][]int{}
	curMinDiff := max - min
	for i, j := 0, 1; j < rangeSize; j++ {
		if r[j] {
			diff := j - i
			if diff < curMinDiff {
				curMinDiff = diff
				result = [][]int{{i - offset, j - offset}}
			} else if diff == curMinDiff {
				result = append(result, []int{i - offset, j - offset})
			}
			i = j
		}
	}

	return result
}

func main() {
	nums := []int{1, 3, 6, 10, 15}
	nums = []int{3, 8, -10, 23, 19, -4, -14, 27}
	nums = []int{4, 2, 1, 3}
	fmt.Println(minimumAbsDifference(nums))
}
