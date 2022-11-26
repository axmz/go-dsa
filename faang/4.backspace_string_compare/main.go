package main

import (
	"fmt"
)

const backspace = '#'

func GetLast(s string) (result string, rest string) {
	if s == "" {
		return s, s
	}

	for l, backspaceCount := len(s)-1, 0; l >= 0; l-- {
		last := s[l]
		if last != backspace {
			if l < backspaceCount {
				return "", ""
			}
			if backspaceCount != 0 {
				backspaceCount--
			} else {
				return string(last), s[0:l]
			}
		} else {
			backspaceCount++
		}
	}

	return "", ""
}

func BackspaceCompare(s string, t string) bool {
	var sLast string
	var tLast string
	var sEnd string = s
	var tEnd string = t

	for {
		sLast, sEnd = GetLast(sEnd)
		tLast, tEnd = GetLast(tEnd)
		if sLast != tLast {
			return false
		}
		if sEnd == "" || tEnd == "" {
			break
		}
	}

	return sEnd == tEnd
}

func main() {
	var s, t string = "bbbextm", "bbb#extm"
	result := BackspaceCompare(s, t)

	fmt.Println("Result: ", result)
}
