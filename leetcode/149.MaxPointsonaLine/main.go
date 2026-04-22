package main

import (
	"fmt"
	"math"
)

func maxPoints(points [][]int) int {
	if len(points) < 3 {
		return len(points)
	}

	max := 0
	for i := 0; i < len(points); i++ {
		m := make(map[float64]int)
		for j := 0; j < len(points); j++ {
			if j != i {
				t := math.Atan2(float64(points[j][0]-points[i][0]), float64(points[j][1]-points[i][1]))
				m[t]++
				if m[t] > max {
					max = m[t]
				}
			}
		}
	}

	return max
}

func maxPoints2(points [][]int) int {
	if len(points) < 3 {
		return len(points)
	}

	max := 0
	for i := 0; i < len(points); i++ {
		samePoints := 1
		slopes := make(map[float64]int)

		for j := i + 1; j < len(points); j++ {
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] {
				samePoints++
			} else {
				var slope float64
				if points[i][0] == points[j][0] {
					slope = 1e9 // vertical line
				} else {
					slope = float64(points[j][1]-points[i][1]) / float64(points[j][0]-points[i][0])
				}
				slopes[slope]++
			}
		}

		currentMax := samePoints
		for _, count := range slopes {
			if count+samePoints > currentMax {
				currentMax = count + samePoints
			}
		}

		if currentMax > max {
			max = currentMax
		}
	}

	return max
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
