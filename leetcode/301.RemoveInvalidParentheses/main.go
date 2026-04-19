package main

import "fmt"

func reverse(s string) string {
	// (()abc)
	// )cba)((
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

// couldn't solve this.
// an interesting technique is to reverse the string and check for the other parentheses
func removeInvalidParentheses(s string) []string {
	var res []string

	var dfs func(s string, start, lastRemove int, par [2]byte)
	dfs = func(s string, start, lastRemove int, par [2]byte) {
		count := 0

		for i := start; i < len(s); i++ {
			if s[i] == par[0] {
				count++
			}
			if s[i] == par[1] {
				count--
			}

			if count >= 0 {
				continue
			}

			// too many closing parentheses → try removing one
			for j := lastRemove; j <= i; j++ {
				if s[j] == par[1] && (j == lastRemove || s[j-1] != par[1]) {
					dfs(s[:j]+s[j+1:], i, j, par)
				}
			}
			return
		}

		// no extra closing parentheses → reverse and check the other side
		reversed := reverse(s)

		if par[0] == '(' {
			dfs(reversed, 0, 0, [2]byte{')', '('})
		} else {
			res = append(res, reversed)
		}
	}

	dfs(s, 0, 0, [2]byte{'(', ')'})
	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
