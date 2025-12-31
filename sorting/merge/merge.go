package merge

// Sort performs merge sort and returns a newly allocated sorted slice.
// A single scratch buffer is reused across the recursion to avoid
// allocating at every merge step (two allocations total: copy+buffer).
func Sort(slc []int) []int {
	s := append([]int{}, slc...)
	if len(s) <= 1 {
		return s
	}
	buf := make([]int, len(s))
	mergeSort(s, buf)
	return s
}

// SortAlloc performs merge sort but allocates at each merge step.
// Clearer but less efficient.
func SortAlloc(slc []int) []int {
	s := append([]int{}, slc...)
	if len(s) <= 1 {
		return s
	}
	return splitAlloc(s)
}

func splitAlloc(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	mid := len(s) / 2
	left := splitAlloc(s[:mid])
	right := splitAlloc(s[mid:])
	return mergeAlloc(left, right)
}

func mergeAlloc(left, right []int) []int {
	ll := len(left)
	lr := len(right)
	merged := make([]int, ll+lr)
	i, j, k := 0, 0, 0

	for i < ll && j < lr {
		if left[i] <= right[j] {
			merged[k] = left[i]
			i++
		} else {
			merged[k] = right[j]
			j++
		}
		k++
	}

	if i < ll {
		copy(merged[k:], left[i:])
	}
	if j < lr {
		copy(merged[k:], right[j:])
	}

	return merged
}

func mergeSort(dst, buf []int) {
	if len(dst) <= 1 {
		return
	}

	mid := len(dst) / 2
	mergeSort(dst[:mid], buf[:mid])
	mergeSort(dst[mid:], buf[mid:])

	// Merge the two sorted halves into the buffer, then copy back.
	i, j, k := 0, mid, 0
	for i < mid && j < len(dst) {
		if dst[i] <= dst[j] {
			buf[k] = dst[i]
			i++
		} else {
			buf[k] = dst[j]
			j++
		}
		k++
	}

	if i < mid {
		copy(buf[k:], dst[i:mid])
	} else if j < len(dst) {
		copy(buf[k:], dst[j:])
	}

	copy(dst, buf[:len(dst)])
}
