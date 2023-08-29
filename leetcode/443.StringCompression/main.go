package main

import (
	"fmt"
	"strconv"
)

// other solutions are more clever
func compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}

	count := 1
	cur := chars[0]
	endIdx := 0
	for i := 1; i < len(chars); i++ {

		if chars[i] == cur && i < len(chars)-1 {
			count++
			continue
		}

		if chars[i] == cur && i == len(chars)-1 {
			count++
		}

		spread := []byte{cur}
		if count > 1 {
			spread = append(spread, []byte(strconv.Itoa(count))...)
		}
		if chars[i] != cur && i == len(chars)-1 {
			spread = append(spread, chars[i])
		}

		_ = append(chars[endIdx:endIdx], spread...)
		endIdx = endIdx + len(spread)
		count = 1
		cur = chars[i]
	}

	chars = chars[:endIdx]
	fmt.Println(string(chars))
	return len(chars)
}

func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c', 'd'}
	// chars := []byte{'a', 'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	// chars := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	// chars := []byte{'a', 'b', 'c', 'c', 'c', 'c', 'c', 'c'}
	res := compress(chars)
	fmt.Println(res)
}
