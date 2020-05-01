package main

import (
	"fmt"
	f "godsa/recursion/factorial"
	fibo "godsa/recursion/fibonacci"
	"time"
)

func main() {
	println("Factorial")
	ifactor := f.Ifactorial(5)
	rfactor := f.Rfactorial(5)
	println("ifactorial", ifactor)
	println("rfactorial", rfactor)

	println("Fibonacci")

	ifibo1 := fibo.Ifibonacci1(10)
	ifibo2 := fibo.Ifibonacci2(10)
	ifibo3 := fibo.Ifibonacci3(10)

	start := time.Now()
	rfibo := fibo.Rfibonacci(10)
	fmt.Println(time.Since(start))

	fmt.Printf("|%-10s|%5d|\n", "ifibo1", ifibo1)
	fmt.Printf("|%-10s|%5d|\n", "ifibo2", ifibo2)
	fmt.Printf("|%-10s|%5d|\n", "ifibo3", ifibo3)
	fmt.Printf("|%-10s|%5d|\n", "rfibo", rfibo)
}
