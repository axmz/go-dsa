package main

import "fmt"

func fizzBuzz(n int) []string {
	res := make([]string, n)

	for i := 1; i <= n; i++ {
		s := ""
		switch {
		case i%3 == 0 && i%5 == 0:
			s += "FizzBuzz"
		case i%3 == 0:
			s += "Fizz"
		case i%5 == 0:
			s += "Buzz"
		default:
			s += fmt.Sprintf("%d", i)
		}
		res[i-1] = s
	}

	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
