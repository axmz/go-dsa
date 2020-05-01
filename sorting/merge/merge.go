package merge

import (
	"math"
)

// Sort merge sorting
func Sort(slc []int) []int {
	s := append([]int{}, slc...)

	return split(s)
}

func split(s []int) []int {
	var l float64 = float64(len(s))
	if l == 1 {
		return s
	}

	middle := int(math.Round(l / 2))
	left := s[:middle]
	right := s[middle:]

	return merge(split(left), split(right))
}

func merge(left []int, right []int) []int {
	var merged []int = []int{}
	ll := len(left)
	lr := len(right)

	var i, j int = 0, 0
	for i < ll && j < lr {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	m := append(left[i:], right[j:]...)
	merged = append(merged, m...)

	return merged
}
