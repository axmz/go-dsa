package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	originalColor := image[sr][sc]
	r, c := len(image), len(image[0])
	visited := make([][]bool, r)
	for i := range visited {
		visited[i] = make([]bool, c)
	}

	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	stack := make([][2]int, 0)
	stack = append(stack, [2]int{sr, sc})
	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curR, curC := pop[0], pop[1]
		visited[curR][curC] = true
		image[curR][curC] = color
		for _, dir := range directions {
			newR, newC := curR+dir[0], curC+dir[1]
			if newR >= 0 && newR < r && newC >= 0 && newC < c {
				if image[newR][newC] == originalColor {
					if !visited[newR][newC] {
						stack = append(stack, [2]int{newR, newC})
					}
				}
			}
		}
	}

	return image
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
