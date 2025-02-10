package main

import "fmt"

func replaceElements(arr []int) []int {
	max := -1

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] > max {
			temp := arr[i]
			arr[i] = max
			max = temp
		} else {
			arr[i] = max
		}
	}

	return arr
}

func main() {
	nums := []int{17, 18, 5, 4, 6, 1}
	fmt.Println(replaceElements(nums))
}
