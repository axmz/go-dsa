package main

import (
	"fmt"
	"strconv"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	var sb strings.Builder
	if (numerator < 0) != (denominator < 0) {
		sb.WriteByte('-')
	}

	n := int64(numerator)
	d := int64(denominator)
	if n < 0 {
		n = -n
	}
	if d < 0 {
		d = -d
	}

	sb.WriteString(strconv.FormatInt(n/d, 10))
	rem := n % d
	if rem == 0 {
		return sb.String()
	}

	sb.WriteByte('.')
	posMap := make(map[int64]int)

	for rem != 0 {
		if p, ok := posMap[rem]; ok {
			s := sb.String()
			return s[:p] + "(" + s[p:] + ")"
		}
		posMap[rem] = sb.Len()
		rem *= 10
		sb.WriteString(strconv.FormatInt(rem/d, 10))
		rem = rem % d
	}
	return sb.String()
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
