package main

import "fmt"

func abs(a int) int {
	if a < 0 {
		return -1 * a
	} else {
		return a
	}
}

func sameDirection(a, b int) bool {
	return a+b == abs(a)+abs(b)
}

func asteroidCollision(asteroids []int) []int {
	stack := []int{}

	for i := 0; i < len(asteroids); i++ {
		cur := asteroids[i]

		if len(stack) == 0 {
			stack = append(stack, cur)
			continue
		}

		peek := stack[len(stack)-1]

		if yes := sameDirection(peek, cur); yes {
			stack = append(stack, cur)
		} else {
			for len(stack) > 0 {
				if peek < 0 {
					stack = append(stack, cur)
					break
				}

				if abs(peek) > abs(cur) {
					break
				}

				if abs(peek) < abs(cur) {
					stack = stack[:len(stack)-1]

					if len(stack) == 0 {
						stack = append(stack, cur)
						break
					}

					peek = stack[len(stack)-1]
					continue
				}

				if abs(peek) == abs(cur) {
					stack = stack[:len(stack)-1]
					break
				}

			}
		}
	}

	return stack
}

func main() {
	// t := []int{5, 10, -5}
	// t := []int{8, -8}
	// t := []int{10, 2, 1, -10}
	// t := []int{10, 2, 1}
	// t := []int{10, 2, -5}
	// t := []int{-2, -1, 1, 2}
	// t := []int{1, -1, -2, -2}
	t := []int{-2, -1, 1, 2, -20}
	r := asteroidCollision(t)
	fmt.Println(r)
}
