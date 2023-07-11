package main

import (
	"fmt"
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	var prev, next int

	for i := 0; i < len(flowerbed) && n != 0; i++ {
		if i == len(flowerbed)-1 {
			next = 0
		} else {
			next = flowerbed[i+1]
		}

		if prev == 0 && flowerbed[i] == 0 && next == 0 {
			flowerbed[i] = 1
			n--
		}
		prev = flowerbed[i]
	}

	return n == 0
}

func main() {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 1
	// n := 2

	res := canPlaceFlowers(flowerbed, n)
	fmt.Println(res)
}
