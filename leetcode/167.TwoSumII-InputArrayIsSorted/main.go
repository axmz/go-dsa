package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	res := []int{0, 0}
	for l < r {
		if numbers[l]+numbers[r] == target {
			res[0] = l + 1
			res[1] = r + 1
			break
		}

		if numbers[l]+numbers[r] > target {
			r--
		}

		if numbers[l]+numbers[r] < target {
			l++
		}
	}

	return res
}

func main() {
	target := 9
	nums := []int{2, 3, 6, 8, 11, 15}
	fmt.Println(twoSum(nums, target))
}
