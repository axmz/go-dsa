package main

import "fmt"

func titleToNumber(columnTitle string) int {
	result := 0
	base := 26
	l := len(columnTitle)
	for i := 0; i < l; i++ {
		j := l - 1 - i
		b := int(columnTitle[i] - 'A' + 1)
		result += b * pow(base, j)
	}
	return result
}

func pow(a, b int) int {
	result := 1
	for b > 0 {
		result *= a
		b--
	}
	return result
}

func main() {
	fmt.Println(titleToNumber("A"))       // 1
	fmt.Println(titleToNumber("AB"))      // 28
	fmt.Println(titleToNumber("ZY"))      // 701
	fmt.Println(titleToNumber("FXSHRXW")) // 2147483647
}
