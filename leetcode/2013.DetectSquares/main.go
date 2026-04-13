package main

import "fmt"

type DetectSquares struct {
	points map[[2]int]int
}

func Constructor() DetectSquares {
	return DetectSquares{
		points: make(map[[2]int]int),
	}
}

func (this *DetectSquares) Add(point []int) {
	this.points[[2]int{point[0], point[1]}]++
}

func (this *DetectSquares) Count(point []int) int {
	queryX, queryY := point[0], point[1]
	total := 0

	for currentPoint, diagonalCount := range this.points {
		currentX, currentY := currentPoint[0], currentPoint[1]

		//  skip same-line points
		if currentX == queryX {
			continue
		}

		// Skip points that are not on the same diagonal
		// Compute the horizontal distance between the stored point and the query point: |currentX - queryX|
		// Compute the vertical distance between them: |currentY - queryY|
		// If those distances are not equal, the two points do not form a square diagonal
		if abs(currentX-queryX) != abs(currentY-queryY) {
			continue
		}

		// product of necessary points that would complete the square with the current diagonal
		total += diagonalCount * this.points[[2]int{currentX, queryY}] * this.points[[2]int{queryX, currentY}]
	}

	return total
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
