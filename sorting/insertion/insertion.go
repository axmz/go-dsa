package insertion

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// Sort insertion sort
func Sort2(original []int) []int {
	s := append([]int{}, original...)
	l := len(s)

	for i := 1; i < l; i++ {
		j := i - 1
		for j >= 0 && s[j] > s[j+1] {
			temp := s[j+1]
			s[j+1] = s[j]
			s[j] = temp
			j--
		}
	}

	return s
}

// Sort insertion sort
func Sort(original []int) []int {
	slc := append([]int{}, original...)
	l := len(slc)

	for i := 1; i < l; i++ {
		if slc[i] <= slc[0] {
			slc = append([]int{slc[i]}, slc[0:]...)
			slc = remove(slc, i+1)
			continue
		} else if slc[i] >= slc[i-1] {
			continue
		} else {
			for j := 1; j < i; j++ {
				if slc[i] <= slc[j] {
					slc = append(slc[:j], append([]int{slc[i]}, slc[j:]...)...)
					slc = remove(slc, i+1)
					break
				}
			}
		}
	}
	return slc
}
