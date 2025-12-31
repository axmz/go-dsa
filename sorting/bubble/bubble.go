package bubble

func Sort(original []int) []int {
	n := len(original)
	arr := make([]int, n)
	copy(arr, original)

	for n > 1 {
		swapped := false
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
		}
		if !swapped {
			break
		}
		n--
	}

	return arr
}
