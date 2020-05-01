package selection

// Sort selection sorting
func Sort(originalArr []int) []int {
	var arr []int
	arr = append(arr, originalArr...)
	l := len(arr)

	for i := 0; i < l; i++ {
		min := i

		for j := i + 1; j < l; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		temp := arr[min]
		arr[min] = arr[i]
		arr[i] = temp
	}

	return arr
}
