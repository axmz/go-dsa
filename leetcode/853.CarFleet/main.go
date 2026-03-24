package main

import (
	"fmt"
	"sort"
)

func carFleet(target int, position []int, speed []int) int {
	l := len(position)
	if l == 1 {
		return 1
	}
	cars := make([][2]int, l)
	for i := range cars {
		cars[i] = [2]int{position[i], speed[i]}
	}

	sort.Slice(cars, func(i int, j int) bool {
		return cars[i][0] < cars[j][0]
	})

	time := make([]float64, l)
	for i := range time {
		time[i] = float64(target-cars[i][0]) / float64(cars[i][1])
	}

	fleets := 1
	for i := l - 2; i >= 0; i-- {
		if time[i] > time[i+1] {
			fleets++
		} else {
			time[i] = time[i+1]
		}
	}

	return fleets
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
