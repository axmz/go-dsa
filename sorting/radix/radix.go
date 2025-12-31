package radix

// countingSort is the original bucket-based implementation (keeps for reference).
func countingSort(arr []int, exp int) []int {
	count := [10][]int{}
	for i, n := range arr {
		digit := (n / exp) % 10
		count[digit] = append(count[digit], arr[i])
	}

	newArr := make([]int, 0, len(arr))
	for _, v := range count {
		if len(v) > 0 {
			newArr = append(newArr, v...)
		}
	}

	return newArr
}

// countingClever is a low-allocation stable counting pass for one digit.
// It allocates exactly one output slice of size len(arr) and uses a fixed
// small counting array (10 ints) — far fewer allocations than the bucket approach.
func countingClever(arr []int, exp int) []int {
	n := len(arr)
	out := make([]int, n)
	var count [10]int

	// count digits
	for _, v := range arr {
		d := (v / exp) % 10
		count[d]++
	}

	// prefix sums
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// build output from end to preserve stability
	for i := n - 1; i >= 0; i-- {
		v := arr[i]
		d := (v / exp) % 10
		count[d]--
		out[count[d]] = v
	}

	return out
}

func Sort(original []int) []int {
	n := len(original)
	if n == 0 {
		return nil
	}

	arr := make([]int, n)
	copy(arr, original)

	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	for exp := 1; max/exp > 0; exp *= 10 {
		arr = countingClever(arr, exp)
	}

	return arr
}
