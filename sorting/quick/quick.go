package quick

func Sort(slc []int) []int {
	s := append([]int{}, slc...)
	if len(s) <= 1 {
		return s
	}
	quicksort(s, 0, len(s)-1)
	return s
}

func quicksort(a []int, lo, hi int) {
	for lo < hi {
		p := partition(a, lo, hi)
		// recurse on smaller partition first to keep stack shallow
		if p-lo < hi-p {
			quicksort(a, lo, p-1)
			lo = p + 1
		} else {
			quicksort(a, p+1, hi)
			hi = p - 1
		}
	}
}

func partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}
