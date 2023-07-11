package main

import (
	"fmt"
)

// func kidsWithCandies(candies []int, extraCandies int) []bool {
// 	res := make([]bool, len(candies))
// 	max := 0
// 	for i, v := range candies {
// 		if v > max {
// 			max = v
// 		}
// 		candies[i] = v + extraCandies
// 	}

// 	for i := range res {
// 		if candies[i] >= max {
// 			res[i] = true
// 		}
// 	}

//		return res
//	}
func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	max := 0
	for _, v := range candies {
		if v > max {
			max = v
		}
	}

	compare := max - extraCandies

	for i := range res {
		if candies[i] >= compare {
			res[i] = true
		}
	}

	return res
}

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3

	res := kidsWithCandies(candies, extraCandies)
	fmt.Println(res)
}
