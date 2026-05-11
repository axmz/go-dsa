package main

import "fmt"

func calPoints(operations []string) int {
	scores := []int{}
	for _, op := range operations {
		switch op {
		case "C":
			scores = scores[:len(scores)-1]
		case "D":
			scores = append(scores, 2*scores[len(scores)-1])
		case "+":
			scores = append(scores, scores[len(scores)-1]+scores[len(scores)-2])
		default:
			var score int
			fmt.Sscanf(op, "%d", &score)
			scores = append(scores, score)
		}
	}

	sum := 0
	for _, score := range scores {
		sum += score
	}
	return sum
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
