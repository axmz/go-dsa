package main

import "fmt"

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	var count int
	numbers := make([]bool, n)

	for i := 2; i < n; i++ {
		if !numbers[i] {
			count++
			for j := i * 2; j < n; j += i {
				numbers[j] = true
			}
		}
	}
	return count
}

func main() {
	n := 12
	result := countPrimes(n)
	fmt.Println(result)
}
