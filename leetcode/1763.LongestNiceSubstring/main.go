package main

import "fmt"

// XOR with space inverts case
func invert(b byte) byte {
	return b ^ 0x20 // 32 is SPC also it is the difference between 'a' and 'A'
}

func isNice(s string) bool {
	var lower uint32 = 0
	var upper uint32 = 0
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b >= 'a' && b <= 'z' {
			lower |= 1 << (b - 'a') // set bit for lowercase
		} else if b >= 'A' && b <= 'Z' {
			upper |= 1 << (b - 'A') // set bit for uppercase
		}
	}
	return lower == upper // if both bitsets are equal, it's a nice substring
}

func longestNiceSubstring2(s string) string {
	n := len(s)
	l, r := 0, 0
	// Substring iteration technique - goes through all substrings
	for i := 0; i < n; i++ {
		for j := i; j <= n; j++ {
			if isNice(s[i:j]) {
				if j-i > r-l {
					l, r = i, j
				}
			}
		}
	}

	return s[l:r]
}

type NiceBitset struct {
	lower    uint32
	upper    uint32
	badChars uint32
}

func NewNiceBitset(s string) *NiceBitset {
	var lower uint32 = 0
	var upper uint32 = 0
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b >= 'a' && b <= 'z' {
			lower |= 1 << (b - 'a') // set bit for lowercase
		} else if b >= 'A' && b <= 'Z' {
			upper |= 1 << (b - 'A') // set bit for uppercase
		}
	}

	badChars := (lower ^ upper) // precompute bitset of chars that don't have both cases

	return &NiceBitset{
		lower:    lower,
		upper:    upper,
		badChars: badChars,
	}
}

func (nb *NiceBitset) IsNice() bool {
	return nb.badChars == 0
}

func (nb *NiceBitset) isBadChar(b byte) bool {
	if b >= 'a' && b <= 'z' {
		return (nb.badChars & (1 << (b - 'a'))) != 0
	} else if b >= 'A' && b <= 'Z' {
		return (nb.badChars & (1 << (b - 'A'))) != 0
	}
	return false
}

func longestNiceSubstring(s string) string {
	if len(s) < 2 {
		return ""
	}

	nb := NewNiceBitset(s)
	for i := range s { // finds first badChar, this is where one needs to split and deepen the search on the left and right
		b := s[i]
		if nb.isBadChar(b) {
			left := longestNiceSubstring(s[0:i])
			right := longestNiceSubstring(s[i+1:])
			if len(left) >= len(right) {
				return left
			} else {
				return right
			}
		}
	}

	return s // if no badChar found, the whole string is nice
}

func main() {
	s := "YazaAay"
	fmt.Println(longestNiceSubstring(s))
}
