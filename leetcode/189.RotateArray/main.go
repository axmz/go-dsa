package main

import "fmt"

func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	start, count := 0, 0

	for count < l {
		current := start
		for {
			next := (current + k) % l
			nums[next], nums[start] = nums[start], nums[next] // ✅ Correct swap
			current = next
			count++
			if start == current {
				break
			}
		}
		start++
	}
}

func rotate4(nums []int, k int) {
	l := len(nums)
	k = k % l

	start, count := 0, 0
	for count < l {
		current, prev := start, nums[start]
		for {
			next := (current + k) % l
			nums[next], prev = prev, nums[next]
			current = next
			count++
			if start == current {
				break
			}
		}
		start++
	}
}

func rotate2(nums []int, k int) {
	l := len(nums)
	for i := l - k; i > 0; i-- {
		for j := i; j < l; j++ {
			nums[j-1], nums[j] = nums[j], nums[j-1]
		}
	}
}
func rotate3(nums []int, k int) {
	for i := 0; i < k; i++ {
		for j := len(nums) - 1; j > 0; j-- {
			nums[j-1], nums[j] = nums[j], nums[j-1]
		}
	}
}

func main() {
	k := 3
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, k)
	fmt.Println(nums)
}
