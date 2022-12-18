package main

import (
	"fmt"
	"strings"
)

var (
	po = "("
	pc = ")"
)

func minRemoveToMakeValid(s string) string {
	var stack []int
	slice := strings.Split(s, "")

	for i, r := range slice {
		if r == po {
			stack = append(stack, i)
		} else if r == pc && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		} else if r == pc {
			slice[i] = "" // try removing in place
		}
	}

	for _, v := range stack {
		slice[v] = ""
	}

	return strings.Join(slice, "")
}

func main() {

	// s := "a(b(c)d"
	// s := "a)b(c)d"
	s := "))(("
	r := minRemoveToMakeValid(s)
	fmt.Println("Result: ", r)
}
