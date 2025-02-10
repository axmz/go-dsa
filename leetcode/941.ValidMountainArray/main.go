package main

import "fmt"

func validMountainArray(arr []int) bool {
	l := len(arr)
	i := 0

	for i+1 < l && arr[i+1] > arr[i] {
		i++
	}

	if i == 0 || i == l-1 {
		return false
	}

	for i+1 < l && arr[i+1] < arr[i] {
		i++
	}

	return i == l-1
}

func main() {
	nums := []int{2, 1}
	// nums := []int{3, 5, 5}
	// nums := []int{0, 3, 2, 1}
	fmt.Println(validMountainArray(nums))
}
