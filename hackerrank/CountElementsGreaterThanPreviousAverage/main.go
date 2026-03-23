package main

import "fmt"

func countResponseTimeRegressions(responseTimes []int32) int32 {
	l := len(responseTimes)
	if l <= 1 {
		return 0
	}

	var count int32 = 0
	var sum float64 = float64(responseTimes[0])
	var avg float64 = sum

	for i := 1; i < l; i++ {
		if avg-float64(responseTimes[i]) < 0 {
			count++
		}
		sum += float64(responseTimes[i])
		avg = sum / float64((int32(i) + 1))
	}

	return count
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
