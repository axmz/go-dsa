package main

import "fmt"

func shipWithinDays(weights []int, days int) int {
	maxLoad := 0
	totalLoad := 0
	for _, w := range weights {
		if w > maxLoad {
			maxLoad = w
		}
		totalLoad += w
	}

	canShip := func(load int) bool {
		daysNeeded := 1
		currentLoad := 0
		for _, w := range weights {
			if currentLoad+w > load {
				daysNeeded++
				currentLoad = 0
			}
			currentLoad += w
		}
		return daysNeeded <= days
	}

	l, r := maxLoad, totalLoad
	for l < r {
		mid := (r-l)/2 + l
		if canShip(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
