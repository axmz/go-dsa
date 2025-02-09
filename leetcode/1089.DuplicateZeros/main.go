package main

import "fmt"

func duplicateZeros(arr []int) {
	posDups := 0
	length := len(arr) - 1

	for left := 0; left <= length-posDups; left++ {
		if arr[left] == 0 {
			if left == length-posDups {
				arr[length] = 0
				length--
				break
			}
			posDups++
		}
	}

	right := length - posDups

	for i := right; i >= 0; i-- {
		if arr[i] == 0 {
			arr[i+posDups] = 0
			posDups--
			arr[i+posDups] = 0
		} else {
			arr[i+posDups] = arr[i]
		}
	}
}

func main() {
	// arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	arr := []int{0, 0, 0, 1, 0, 0, 0}
	duplicateZeros(arr)
	fmt.Println(arr)
}
