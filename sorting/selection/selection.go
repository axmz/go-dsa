package selection

func Sort(original []int) []int {
	n := len(original)
	arr := make([]int, n)
	copy(arr, original)

	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}

	return arr
}
