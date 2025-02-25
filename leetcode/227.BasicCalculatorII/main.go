package main

import (
	"fmt"
)

func calculate(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}

	var currentNumber, lastNumber, result int
	var sign byte = '+'

	for i := 0; i < length; i++ {
		cur := s[i]

		if cur >= '0' && cur <= '9' {
			currentNumber = currentNumber*10 + int(cur-'0')
		} else if cur == ' ' {
			continue
		} else if cur == '+' || cur == '-' || cur == '*' || cur == '/' {
			if sign == '+' || sign == '-' {
				result += lastNumber
				if sign == '+' {
					lastNumber = currentNumber
				} else {
					lastNumber = -currentNumber
				}
			} else if sign == '*' {
				lastNumber = lastNumber * currentNumber
			} else if sign == '/' {
				lastNumber = lastNumber / currentNumber
			}

			sign = cur
			currentNumber = 0
		}
	}

	// Handle the last number after the loop
	if sign == '+' {
		result += lastNumber + currentNumber
	} else if sign == '-' {
		result += lastNumber - currentNumber
	} else if sign == '*' {
		result += lastNumber * currentNumber
	} else if sign == '/' {
		result += lastNumber / currentNumber
	}

	return result
}

func main() {
	s := " 3+5 / 2 "
	fmt.Println(s, calculate(s))
}
