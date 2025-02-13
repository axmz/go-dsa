package main

import "fmt"

func getRow(rowIndex int) []int {
	prev := []int{}
	for i := 0; i <= rowIndex; i++ {
		arr := make([]int, i+1)
		arr[0], arr[len(arr)-1] = 1, 1

		if i > 1 {
			for k, l := 1, 2; l < len(arr); k, l = k+1, l+1 {
				arr[k] = prev[k-1] + prev[k]
			}
		}
		prev = arr
	}

	return prev
}

func main() {
	x := 5
	fmt.Println(getRow(x))
}
