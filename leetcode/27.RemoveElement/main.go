package main

import "fmt"

func removeElement2(nums []int, val int) int {
	count := 0
	for i, j := 0, len(nums)-1; i <= j; {
		if nums[i] == val {
			nums[i] = nums[j]
			nums[j] = -1
			if nums[i] != val {
				i++
				count++
			}
			j--
		} else {
			i++
			count++
		}
	}
	return count
}

func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	// nums := []int{3, 2, 2, 3}
	// val := 3
	fmt.Println(removeElement(nums, val))
	fmt.Println(nums)
}
