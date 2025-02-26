package main

import "fmt"

type deque struct {
	indexes []int
}

func (d *deque) push(i int) {
	d.indexes = append(d.indexes, i)
}

func (d *deque) getFirst() int {
	return d.indexes[0]
}

func (d *deque) popFirst() {
	d.indexes = d.indexes[1:]
}

func (d *deque) getLast() int {
	return d.indexes[len(d.indexes)-1]
}

func (d *deque) popLast() {
	d.indexes = d.indexes[:len(d.indexes)-1]
}

func (d *deque) empty() bool {
	return len(d.indexes) == 0
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < k || k == 0 {
		return make([]int, 0)
	} else if k == 1 {
		return nums
	}

	var (
		// len(nums)-k+1 this is the number of windows.
		// If input has length 3 and k is 1, we'll have 3 - 1 + 1 = 3 windows.
		res = make([]int, len(nums)-k+1)
		dq  = &deque{}
	)

	for i := range nums {
		// we pop the first element if it's outside of the current window.
		if !dq.empty() && (i-k == dq.getFirst()) {
			dq.popFirst()
		}

		// we pop all the elements that are smaller than the current one.
		for !dq.empty() && nums[dq.getLast()] < nums[i] {
			dq.popLast()
		}

		// we push the current one to the window.
		dq.push(i)

		// if we reached at least the first window end.
		if i >= k-1 {
			// we add the current result that is the first in the deque.
			res[i-k+1] = nums[dq.getFirst()]
		}
	}

	return res
}

func maxSlidingWindow2(nums []int, k int) []int {
	l := len(nums)

	if l == 0 || k == 0 {
		return []int{}
	}

	d := []int{}
	res := make([]int, 0, l-k+1)

	for i := 0; i < k; i++ {
		for len(d) > 0 && nums[i] >= nums[d[len(d)-1]] {
			d = d[:len(d)-1]
		}

		d = append(d, i)
	}

	res = append(res, nums[d[0]])

	for i := k; i < l; i++ {
		if d[0] == i-k {
			d = d[1:]
		}

		for len(d) > 0 && nums[i] >= nums[d[len(d)-1]] {
			d = d[:len(d)-1]
		}

		d = append(d, i)

		res = append(res, nums[d[0]])
	}

	return res
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
