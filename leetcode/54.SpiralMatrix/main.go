package main

import "fmt"

func spiralOrderGoroutines(matrix [][]int) []int {
	w := len(matrix[0])
	h := len(matrix)
	arr := make([]int, 0, w*h)
	ch := make(chan int)

	move := func(ch chan int) {
		x := 0
		y := 0
		directions := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
		visited := -101

		ch <- matrix[y][x]
		matrix[y][x] = visited
		visitedCount := 1

		for {
			for _, dir := range directions {
				for {
					newX, newY := x+dir[0], y+dir[1]
					if newX >= 0 && newX < w && newY >= 0 && newY < h && matrix[newY][newX] != -101 {
						val := matrix[newY][newX]
						matrix[newY][newX] = visited
						visitedCount++
						x, y = newX, newY
						fmt.Println(x, y)
						ch <- val
					} else {
						break
					}
				}
			}

			if visitedCount == w*h {
				break
			}
		}
		close(ch)
	}

	go move(ch)

	for v := range ch {
		arr = append(arr, v)
	}

	return arr
}

func spiralOrder(matrix [][]int) []int {
	w := len(matrix[0])
	h := len(matrix)
	arr := make([]int, 0, w*h)

	directions := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	visited := -101

	x := 0
	y := 0

	arr = append(arr, matrix[y][x])
	matrix[y][x] = visited
	visitedCount := 1

	for {
		for _, dir := range directions {
			for {
				newX, newY := x+dir[0], y+dir[1]
				if newX >= 0 && newX < w && newY >= 0 && newY < h && matrix[newY][newX] != -101 {
					val := matrix[newY][newX]
					matrix[newY][newX] = visited
					visitedCount++
					x, y = newX, newY
					arr = append(arr, val)
				} else {
					break
				}
			}
		}

		if visitedCount == w*h {
			break
		}
	}

	return arr
}

func main() {
	nums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(nums))
}
