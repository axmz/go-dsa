package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	// define operations as functions
	operations := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	stack := []int{}

	for _, t := range tokens {
		if f, ok := operations[t]; ok {
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			val := f(a, b)
			stack = append(stack, val)
		} else {
			// better convert here once than twice above
			conv, _ := strconv.Atoi(t)
			stack = append(stack, conv)
		}
	}

	return stack[0]
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
