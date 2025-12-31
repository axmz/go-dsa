package insertion

func Sort(original []int) []int {
	n := len(original)
	arr := make([]int, n)
	copy(arr, original)

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j] // or swap
			j--
		}
		arr[j+1] = key
	}

	return arr
}
