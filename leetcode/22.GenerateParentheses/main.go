package main

import "fmt"

func generateParenthesis(n int) []string {
	result := []string{}

	var backtrack func(state string, open int, close int, max int)
	backtrack = func(state string, open int, close int, max int) {
		if len(state) == max*2 {
			result = append(result, state)
			return
		}

		if open < max {
			backtrack(state+"(", open+1, close, max)
		}
		if close < open {
			backtrack(state+")", open, close+1, max)
		}
	}

	backtrack("", 0, 0, n)
	return result
}

func main() {
	x := 4
	fmt.Println(generateParenthesis(x))
}
