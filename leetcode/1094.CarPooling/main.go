package main

import (
	"fmt"
	"sort"
)

func carPooling(trips [][]int, capacity int) bool {
	passengers := 0
	start := 1
	end := 2
	passenger_change := make([][2]int, 0, len(trips)*2)
	for _, trip := range trips {
		passenger_change = append(passenger_change, [2]int{trip[start], trip[passengers]})
		passenger_change = append(passenger_change, [2]int{trip[end], -trip[passengers]})
	}

	sort.Slice(passenger_change, func(i, j int) bool {
		if passenger_change[i][0] == passenger_change[j][0] {
			return passenger_change[i][1] > passenger_change[j][1]
		}
		return passenger_change[i][0] < passenger_change[j][0]
	})

	for _, change := range passenger_change {
		passengers += change[1]
		if passengers > capacity {
			return false
		}
	}

	return true
}
func carPooling(trips [][]int, capacity int) bool {
	passengers := 0
	start := 1
	end := 2
	passenger_change := make([][2]int, 0, len(trips)*2)
	for _, trip := range trips {
		passenger_change = append(passenger_change, [2]int{trip[start], trip[passengers]})
		passenger_change = append(passenger_change, [2]int{trip[end], -trip[passengers]})
	}

	sort.Slice(passenger_change, func(i, j int) bool {
		if passenger_change[i][0] == passenger_change[j][0] {
			return passenger_change[i][1] < passenger_change[j][1]
		}
		return passenger_change[i][0] < passenger_change[j][0]
	})

	for _, change := range passenger_change {
		passengers += change[1]
		if passengers > capacity {
			return false
		}
	}

	return true
}

func main() {
	trips := [][]int{{2, 1, 5}, {3, 3, 7}}
	capacity := 5
	fmt.Println(carPooling(trips, capacity))
}
