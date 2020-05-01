package fibonacci

import (
	"log"
	"time"
)

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

// Ifibonacci1 iterable fibo without pointers
func Ifibonacci1(n int) int {
	// defer duration(track("ifibo1"))

	var result []int = []int{0, 1}

	for i := 1; i < n; i++ {
		l := len(result)
		next := result[l-1] + result[l-2]
		result = append(result, next)
	}

	return result[n]
}

// Ifibonacci2 iterable fibo with pointers
func Ifibonacci2(n int) int {
	// defer duration(track("ifibo2"))

	var result *[]int = &[]int{0, 1}

	for i := 1; i < n; i++ {
		l := len(*result)
		next := (*result)[l-1] + (*result)[l-2]
		*result = append(*result, next)
	}

	return (*result)[n]
}

// Ifibonacci3 iterable fibo with pointers and array instead of slice
func Ifibonacci3(n int) int {
	// defer duration(track("ifibo3"))

	var result *[2]int = &[2]int{0, 1}
	for i := 1; i < n; i++ {
		sum := result[0] + result[1]
		result[0] = result[1]
		result[1] = sum
	}

	return result[1]
}
