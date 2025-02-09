package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	hottest := 0
	res := make([]int, len(temperatures))
	for i := len(temperatures) - 1; i >= 0; i-- {
		if temperatures[i] >= hottest {
			hottest = temperatures[i]
		} else {
			for j := i + 1; ; {
				if temperatures[i] < temperatures[j] {
					res[i] = j - i
					break
				} else {
					j += res[j]
				}
			}
		}
	}

	return res
}

func dailyTemperatures2(temperatures []int) []int {
	stack := [][]int{{0, temperatures[0]}}
	res := make([]int, len(temperatures))

	for i := 1; i < len(temperatures); i++ {
		for len(stack) > 0 {
			pop := stack[len(stack)-1]
			if temperatures[i] > pop[1] {
				res[pop[0]] = i - pop[0]
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, []int{i, temperatures[i]})
	}

	return res
}

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temperatures))
}
