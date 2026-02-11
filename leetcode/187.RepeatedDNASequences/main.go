package main

import (
	"fmt"
)

// Baby Rabin-Karp
func findRepeatedDnaSequences2(s string) []string {
	if len(s) <= 10 {
		return []string{}
	}
	L := 10
	res := []string{}
	seen := make(map[int]bool)
	output := make(map[int]bool)
	translate := map[byte]int{
		'A': 0,
		'C': 1,
		'G': 2,
		'T': 3,
	}

	// Note: This is simplified Rabin-Karp implementation
	// The hash space is enough to avoid collisions for this problem
	// because there are only 4^10 = 1,048,576 possible sequences of length 10
	// Below there is another, more realistic Rabin-Karp implementation

	// rolling hash params
	a := 4
	aL := 1 // a^L-1
	// a^L should in fact be multiplied L-1 times not L times
	// When you build the hash for a substring of length L:
	// The leftmost character is multiplied by a^(L-1)
	// The next by a^(L−2),
	// ...down to the rightmost, which is multiplied by a^0.
	for i := 1; i < L; i++ {
		aL *= a
	}

	// initial hash
	hash := 0
	substr := s[0:L]
	for j := 0; j < L; j++ {
		hash = hash*a + translate[substr[j]]
	}
	seen[hash] = true

	// rolling hash
	for i := L; i < len(s); i++ {
		substr = s[i-L+1 : i+1]
		// update hash
		m := translate[s[i-L]]
		n := translate[s[i]]
		hash = (hash-m*aL)*a + n

		if seen[hash] {
			if output[hash] {
				continue
			}
			output[hash] = true
			res = append(res, substr)
		} else {
			seen[hash] = true
		}
	}

	return res
}

// Rabin-Karp
func findRepeatedDnaSequences(s string) []string {
	if len(s) <= 10 {
		return []string{}
	}
	L := 10
	res := []string{}
	seen := make(map[int]bool)
	output := make(map[int]bool)
	translate := map[byte]int{
		'A': 0,
		'C': 1,
		'G': 2,
		'T': 3,
	}

	// Rabin-Karp more realistic
	// rolling hash params - more robust against collisions
	a := 101           // a prime number
	q := 1_000_000_007 // a large prime number
	aL := 1            // a^L-1
	for i := 1; i < L; i++ {
		aL = (aL * a) % q
	}

	// initial hash
	hash := 0
	substr := s[0:L]
	for j := 0; j < L; j++ {
		hash = (hash*a + translate[substr[j]]) % q
	}
	seen[hash] = true

	// rolling hash
	for i := L; i < len(s); i++ {
		substr = s[i-L+1 : i+1]
		// update hash
		m := translate[s[i-L]]
		n := translate[s[i]]
		hash = ((hash-m*aL)*a + n) % q
		if hash < 0 {
			hash += q
		}

		if seen[hash] {
			if output[hash] {
				continue
			}
			output[hash] = true
			res = append(res, substr)
		} else {
			seen[hash] = true
		}
	}

	return res
}

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
}
