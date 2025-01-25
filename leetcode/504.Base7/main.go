package main

import (
	"fmt"
)

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	rest := num
	if num < 0 {
		rest *= -1
	}

	res := ""
	for rest != 0 {
		m := rest % 7
		rest = rest / 7
		res = string(m+'0') + res
	}

	if num < 0 {
		res = "-" + res
	}

	return res
}

func main() {
	// num := 102
	// num := -10
	num := -7
	fmt.Println(convertToBase7(num))
}
