package main

import "fmt"

func hash(s string) string {
	res := make([]byte, len(s))
	f := s[0]
	for i, v := range s {
		res[i] = (byte(v) - f + 26) % 26
	}

	return string(res)
}

func groupStrings(strings []string) [][]string {
	m := map[string][]string{}
	for i, s := range strings {
		h := hash(s)
		m[h] = append(m[h], strings[i])
	}

	res := [][]string{}
	for _, v := range m {
		res = append(res, v)
	}

	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums, (-2+26)%26)
}
