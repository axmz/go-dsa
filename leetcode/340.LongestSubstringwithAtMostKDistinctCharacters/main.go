package main

import "fmt"

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if k == 0 || len(s) == 0 {
		return 0
	}

	l := 0
	r := 0
	countK := 0
	max := 0
	m := make(map[byte]int)

	add := func(r int) {
		sym := s[r] - 'a'
		if m[sym] == 0 {
			m[sym]++
			countK++
		} else {
			m[sym]++
		}
	}

	rem := func(l int) {
		sym := s[l] - 'a'
		m[sym]--
		if m[sym] == 0 {
			countK--
		}
	}

	for r < len(s) {
		add(r)
		if countK <= k {
			if r-l > max {
				max = r - l
			}
		} else {
			rem(l)
			l++
		}
		r++
	}

	return max + 1
}

func main() {
	s := "eceba"
	k := 2
	fmt.Println(s, k, lengthOfLongestSubstringKDistinct(s, k))
	s = "aa"
	k = 1
	fmt.Println(s, k, lengthOfLongestSubstringKDistinct(s, k))
	s = "ab"
	k = 1
	fmt.Println(s, k, lengthOfLongestSubstringKDistinct(s, k))
}
