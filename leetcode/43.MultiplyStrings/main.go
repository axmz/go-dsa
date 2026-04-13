package main

import (
	"fmt"
)

func multiply(num1 string, num2 string) string {
	b1 := []byte(num1)
	b2 := []byte(num2)

	l1 := len(b1)
	l2 := len(b2)
	result := make([]byte, l1+l2)

	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			digit1 := b1[i] - '0'
			digit2 := b2[j] - '0'
			product := digit1*digit2 + result[i+j+1]
			result[i+j+1] = product % 10
			result[i+j] += product / 10
		}
	}

	// skip leading zeros
	i := 0
	for i < len(result)-1 && result[i] == 0 {
		i++
	}

	// convert result to string
	for k := range result {
		result[k] += '0'
	}

	return string(result[i:])
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
