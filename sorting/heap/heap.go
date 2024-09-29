package heap

func heapifyMin(nums []int, i int) {
	l := len(nums)
	left := 2*i + 1
	right := 2*i + 2
	min := i

	if left < l && nums[left] < nums[min] {
		min = left
	}

	if right < l && nums[right] < nums[min] {
		min = right
	}

	// Swap and continue heapifying if the min is not the current node
	if min != i {
		nums[i], nums[min] = nums[min], nums[i]
		heapifyMin(nums, min)
	}
}

func heapifyMax(nums []int, i int) {
	l := len(nums)
	left := 2*i + 1
	right := 2*i + 2
	max := i

	if right < l && nums[right] > nums[max] {
		max = right
	}

	if left < l && nums[left] > nums[max] {
		max = left
	}

	if max != i {
		nums[i], nums[max] = nums[max], nums[i]
		heapifyMax(nums, max)
	}
}

func Sort(nums []int) []int {
	c := make([]int, len(nums))
	copy(c, nums)

	for i := len(c)/2 - 1; i >= 0; i-- {
		heapifyMax(c, i)
	}

	sorted := c

	for i := len(c) - 1; i > 0; i-- {
		c[0], c[i] = c[i], c[0]
		c = c[:i]
		heapifyMax(c[:i], 0)
	}

	return sorted
}

func SortDesc(nums []int) []int {
	c := make([]int, len(nums))
	copy(c, nums)

	for i := len(c)/2 - 1; i >= 0; i-- {
		heapifyMin(c, i)
	}

	sorted := c

	for i := len(c) - 1; i > 0; i-- {
		c[0], c[i] = c[i], c[0]
		c = c[:i]
		heapifyMin(c[:i], 0)
	}

	return sorted
}
