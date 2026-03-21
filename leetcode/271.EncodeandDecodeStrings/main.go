package main

import (
	"fmt"
	"strings"
)

type Codec struct {
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	sb := strings.Builder{}
	for _, s := range strs {
		l := len(s)
		sb.WriteByte(byte(l))
		if l == 0 {
			sb.WriteByte(byte(0))
			continue
		}

		for i := 0; i < l; i++ {
			sb.WriteByte(s[i])
		}
	}

	return sb.String()
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	res := []string{}

	for i := 0; i < len(strs); {
		l := int(strs[i])
		s := string(strs[i+1 : i+l+1])
		res = append(res, s)
		if l == 0 {
			i += l + 2
		} else {
			i += l + 1
		}
	}

	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
