package main

import "fmt"

func flipAndInvertImage(image [][]int) [][]int {
	for _, r := range image {
		for i := 0; i*2 < len(r); i++ {
			r[i], r[len(r)-1-i] = r[len(r)-1-i]^1, r[i]^1
		}
	}

	return image
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
