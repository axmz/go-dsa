package main

import "fmt"

func convertToTitle(columnNumber int) string {
	res := []byte{}
	for columnNumber > 0 {
		columnNumber--
		res = append([]byte{byte(columnNumber%26) + 'A'}, res...)
		columnNumber /= 26
	}
	return string(res)
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
