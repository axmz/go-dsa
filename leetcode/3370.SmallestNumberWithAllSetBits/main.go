package main

import "fmt"

// O(log(log(n))) -- I would say it is O(1) cause it is constant
// But if you look otherwise you can say that n represented in bit takes O(log(n))
// And to get smallestNumber you need another O(log(n))
func smallestNumber(n int) int {
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	return n
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
