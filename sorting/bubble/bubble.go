package bubble

// Sort bubble sort
func Sort(originalSlice []int) []int {
	var slc = make([]int, len(originalSlice))
	copy(slc, originalSlice)
	l := len(slc)

	for i := 1; i < l; l-- {
		for j := 1; j < l; j++ {
			if slc[j-1] > slc[j] {
				temp := slc[j-1]
				slc[j-1] = slc[j]
				slc[j] = temp
			}
		}
	}

	return slc
}
