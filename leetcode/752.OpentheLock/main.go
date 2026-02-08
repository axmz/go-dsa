package main

import (
	"fmt"
)

type State struct {
	code [4]byte
	step int
}

func (s State) Up(i int) State {
	b := s.code
	digit := int(b[i] - '0')
	newDigit := (digit + 1) % 10
	b[i] = byte(newDigit) + '0'
	return State{b, s.step + 1}
}

func (s State) Down(i int) State {
	b := s.code
	digit := int(b[i] - '0')
	newDigit := (digit + 9) % 10
	b[i] = byte(newDigit) + '0'
	return State{b, s.step + 1}
}

func openLock(deadends []string, target string) int {
	dead := make(map[[4]byte]bool)
	for _, d := range deadends {
		var arr [4]byte
		copy(arr[:], d)
		dead[arr] = true
	}
	var start [4]byte
	copy(start[:], "0000")
	if dead[start] {
		return -1
	}
	visited := make(map[[4]byte]bool)
	q := make([]State, 0)
	q = append(q, State{start, 0})
	visited[start] = true
	targetArr := [4]byte{}
	copy(targetArr[:], target)

	for len(q) > 0 {
		state := q[0]
		q = q[1:]
		if state.code == targetArr {
			return state.step
		}
		for i := 0; i < 4; i++ {
			up := state.Up(i)
			if !dead[up.code] && !visited[up.code] {
				visited[up.code] = true
				q = append(q, up)
			}
			down := state.Down(i)
			if !dead[down.code] && !visited[down.code] {
				visited[down.code] = true
				q = append(q, down)
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
}
