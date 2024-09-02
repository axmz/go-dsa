package main

import "fmt"

func canVisitAllRooms(rooms [][]int) bool {
	seen := make([]bool, len(rooms))
	stack := []int{0}
	var pop int

	for len(stack) > 0 {
		pop, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if !seen[pop] {
			seen[pop] = true
			stack = append(stack, rooms[pop]...)
		}
	}

	count := 0
	for _, v := range seen {
		if v {
			count++
		}
	}

	return count == len(rooms)
}

func main() {
	// rooms := [][]int{{1}, {2}, {3}, {}}
	// rooms := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	rooms := [][]int{{0, 0}}
	res := canVisitAllRooms(rooms)
	fmt.Println(res)
}
