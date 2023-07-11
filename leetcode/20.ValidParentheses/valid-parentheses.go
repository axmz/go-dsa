package main

func isValid(s string) bool {
	if len(s) == 1 {
		return false
	}
	symbols := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}

	l := len(s)

	stack := []string{}

	for i := 0; i < l; i++ {
		sym := string(s[i])
		if sym == "(" || sym == "[" || sym == "{" {
			stack = append(stack, sym)
		} else {
			if len(stack) == 0 {
				return false
			}
			if len(stack) > 0 {
				top := string(stack[len(stack)-1])
				if sym == symbols[top] {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		}
	}
	isEmpty := true
	if len(stack) > 0 {
		isEmpty = false
	}
	return isEmpty
}
