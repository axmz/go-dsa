package main

import (
	"fmt"
	"math"
)

func minEatingSpeed(piles []int, h int) int {
	hoursPerPile := float64(h) / float64(len(piles))
	hoursPerPile = math.Floor(hoursPerPile)
	// min max optimizations
	max := 0
	for _, v := range piles {
		if v > max {
			max = v
		}
	}
	max = int(math.Ceil(float64(max) / hoursPerPile))

	min := 1
	for _, v := range piles {
		if v < min {
			min = v
		}
	}
	min = int(math.Floor(float64(min) / hoursPerPile))
	if min < 1 {
		min = 1
	}

	l, r := min, max

	for l < r {
		speed := (l + r) / 2
		totalHours := 0
		for i := 0; i < len(piles); i++ {
			if totalHours > h {
				break
			}
			totalHours += int(math.Ceil(float64(piles[i]) / float64(speed)))
		}

		if totalHours > h {
			l = speed + 1
		} else {
			r = speed
		}
	}

	return l
}

func main() {
	// piles := []int{3, 6, 7, 11}
	// h := 8
	piles := []int{1000000000, 1000000000}
	h := 3
	fmt.Println(minEatingSpeed(piles, h))
}
