package main

import "fmt"

/**
 * // This is the robot's control interface.
 * // You should not implement it, or speculate about its implementation
 * type Robot struct {
 * }
 *
 * // Returns true if the cell in front is open and robot moves into the cell.
 * // Returns false if the cell in front is blocked and robot stays in the current cell.
 * func (robot *Robot) Move() bool {}
 *
 * // Robot will stay in the same cell after calling TurnLeft/TurnRight.
 * // Each turn will be 90 degrees.
 * func (robot *Robot) TurnLeft() {}
 * func (robot *Robot) TurnRight() {}
 *
 * // Clean the current cell.
 * func (robot *Robot) Clean() {}
 */

type Robot struct {
}

func (robot *Robot) Move() bool {
	fmt.Println("move")
	return true
}

func (robot *Robot) TurnLeft() {
	fmt.Println("left")
}

func (robot *Robot) TurnRight() {
	fmt.Println("right")
}

func (robot *Robot) Clean() {
	fmt.Println("clean")
}

type Set map[[2]int]bool

func (s *Set) Add(r, c int) {
	(*s)[[2]int{r, c}] = true
}

func (s *Set) Contains(r, c int) bool {
	return (*s)[[2]int{r, c}]
}

func cleanRoom(robot *Robot) {
	visited := make(Set)
	directions := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var r, c, currentDirection int

	moveBack := func() {
		robot.TurnLeft()
		robot.TurnLeft()
		robot.Move()
		robot.TurnLeft()
		robot.TurnLeft()
	}

	var backtrack func(r, c, dir int)
	backtrack = func(r, c, dir int) {
		visited.Add(r, c)
		robot.Clean()

		for i := range directions {
			nextDir := (dir + i) % len(directions)
			nextR := r + directions[nextDir][0]
			nextC := c + directions[nextDir][1]

			if !visited.Contains(nextR, nextC) && robot.Move() {
				backtrack(nextR, nextC, nextDir)
			}

			robot.TurnRight()
		}

		moveBack()
	}

	backtrack(r, c, currentDirection)
}

// func cleanRoom1(robot *Robot) {
// 	m := make(map[int]map[int]bool)
// 	directions := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
// 	var r, c, currentDirection int

// 	createMapEntry := func(r, c int) {
// 		if _, exist := m[r]; !exist {
// 			m[r] = map[int]bool{}
// 		}
// 	}

// 	var backtrack func(r, c, dir int)
// 	backtrack = func(r, c, dir int) {
// 		createMapEntry(r, c)
// 		robot.Clean()
// 		m[r][c] = true

// 		for i := 0; i < len(directions); i++ {
// 			newDir := dir
// 			newDir += i
// 			newDir %= len(directions)

// 			newR := r + directions[newDir][0]
// 			newC := c + directions[newDir][1]

// 			var visited, exist bool
// 			visited, exist = m[newR][newC]

// 			if !exist {
// 				createMapEntry(newR, newC)
// 				if !robot.Move() {
// 					m[newR][newC] = true
// 				} else {
// 					backtrack(newR, newC, newDir)
// 					robot.TurnLeft()
// 					robot.TurnLeft()
// 				}
// 			}

// 			if exist {
// 				if !visited && robot.Move() {
// 					backtrack(newR, newC, newDir)
// 					robot.TurnLeft()
// 					robot.TurnLeft()
// 				} else {
// 					m[newR][newC] = true
// 				}
// 			}

// 			robot.TurnRight()
// 		}

// 		robot.TurnLeft()
// 		robot.TurnLeft()
// 		robot.Move()
// 	}

// 	backtrack(r, c, currentDirection)
// }

func main() {
	r := new(Robot)
	cleanRoom(r)
}
