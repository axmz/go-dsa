package main

import "fmt"

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	minDist := len(wordsDict) + 1 // instead of math.MaxInt32
	w1, w2 := -1, -1
	for i, word := range wordsDict {
		if word == word1 {
			w1 = i
		} else if word == word2 {
			w2 = i
		}
		if w1 != -1 && w2 != -1 {
			dist := w2 - w1
			if dist < 0 {
				dist = -dist
			}
			if dist < minDist {
				minDist = dist
			}
		}
	}
	return minDist
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
