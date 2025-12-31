package counting

func Sort(original []int) []int {
	n := len(original)
	arr := make([]int, n)
	copy(arr, original)

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
	r := make([]int, rangeSize)
	offset := -min

	for _, v := range arr {
		r[v+offset]++
	}
	result := []int{}
	for i, count := range r {
		for count > 0 {
			result = append(result, i-offset)
			count--
		}
	}

	return result
}
