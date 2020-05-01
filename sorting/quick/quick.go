package quick

// Sort quick sorting
func Sort(slc []int) []int {
	s := append([]int{}, slc...)
	pivot(s)
	return s
}

func pivot(s []int) []int {
	if len(s) < 2 {
		return s
	}

	l := len(s)
	p := s[l-1]
	j := l - 1
	for i := l - 2; i >= 0; i-- {
		j = l - 1
		c := s[i]
		p = s[j]

		if p < c {
			s[j] = c
			s[i] = p
			j--
		}
	}

	left := s[:j]
	right := s[j+1:]

	return merge(pivot(left), pivot(right), p)
}

func merge(l, r []int, p int) []int {
	merged := []int{}
	merged = append(merged, l...)
	merged = append(merged, p)
	merged = append(merged, r...)
	return merged
}
