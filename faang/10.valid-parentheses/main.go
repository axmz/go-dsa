package main

import "fmt"

var p = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func Last[E any](s []E) (E, bool) {
	if len(s) == 0 {
		var zero E
		return zero, false
	}
	return s[len(s)-1], true
}

func isValid(s string) bool {
	var stack []rune

	for _, r := range s {
		if val, ok := p[r]; ok {
			if l, ok := Last(stack); ok && l == val { // return early
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, r)
		}
	}

	if len(stack) > 0 { // return len(stack) == 0
		return false
	} else {
		return true
	}
}

func main() {
	s := isValid("]")
	fmt.Println("Result: ", s)
}
