package main

import "fmt"

func generate(numRows int) [][]int {
	res := [][]int{}
	for i := 0; i < numRows; i++ {
		arr := make([]int, i+1)
		arr[0], arr[len(arr)-1] = 1, 1
		res = append(res, arr)

		if i > 1 {
			for k, l := 1, 2; l < len(arr); k, l = k+1, l+1 {
				arr[k] = res[i-1][k-1] + res[i-1][k]
			}
		}
	}

	return res
}

func main() {
	x := 5
	fmt.Println(generate(x))
}
