package main

import "fmt"

func countVowelPermutation(n int) int {
	m := 1000000007
	a, e, i, o, u := 1, 1, 1, 1, 1

	for j := 1; j < n; j++ {
		a, e, i, o, u = (e+i+u)%m, (a+i)%m, (e+o)%m, i%m, (i+o)%m
	}

	return (a + e + i + o + u) % m
}

func countVowelPermutation3(n int) int {
	dp := make([][5]int, n)
	for i := 0; i < 5; i++ {
		dp[0][i] = 1
	}

	// Each vowel 'a' may only be followed by an 'e'.	 			e
	// Each vowel 'e' may only be followed by an 'a' or an 'i'.		a, i
	// Each vowel 'i' may not be followed by another 'i'. // 		a, e, o, u
	// Each vowel 'o' may only be followed by an 'i' or a 'u'.		i, u
	// Each vowel 'u' may only be followed by an 'a'.				a
	// a -> e, i, u
	// e -> a, i
	// i -> e, o
	// o -> i
	// u -> i, o
	for i := 1; i < n; i++ {
		dp[i][0] = (dp[i-1][1] + dp[i-1][2] + dp[i-1][4]) % 1000000007
		dp[i][1] = (dp[i-1][0] + dp[i-1][2]) % 1000000007
		dp[i][2] = (dp[i-1][1] + dp[i-1][3]) % 1000000007
		dp[i][3] = dp[i-1][2] % 1000000007
		dp[i][4] = (dp[i-1][2] + dp[i-1][3]) % 1000000007
	}

	count := 0
	for i := 0; i < 5; i++ {
		count = (count + dp[n-1][i]) % 1000000007
	}

	return count
}

func countVowelPermutation2(n int) int {
	// ints instead of bytes to avoid map[[2]interface{}]int
	chars := []int{'a', 'e', 'i', 'o', 'u'}
	memo := make(map[[2]int]int)
	p := 1000000007

	if n == 1 {
		return 5
	}

	// state: prev, i,
	var backtrack func(prev int, i int) int
	backtrack = func(prev int, i int) int {
		if i == n {
			return 1
		}
		if val, ok := memo[[2]int{prev, i}]; ok {
			return val
		}
		count := 0
		for _, c := range chars {
			// Each vowel 'a' may only be followed by an 'e'.
			// Each vowel 'e' may only be followed by an 'a' or an 'i'.
			// Each vowel 'i' may not be followed by another 'i'.
			// Each vowel 'o' may only be followed by an 'i' or a 'u'.
			// Each vowel 'u' may only be followed by an 'a'.
			if (prev == 'a' && c == 'e') ||
				(prev == 'e' && (c == 'a' || c == 'i')) ||
				(prev == 'i' && c != 'i') ||
				(prev == 'o' && (c == 'i' || c == 'u')) ||
				(prev == 'u' && c == 'a') {
				count = (count + backtrack(c, i+1)) % p
			}
		}
		memo[[2]int{prev, i}] = count
		return count
	}

	count := 0
	for _, c := range chars {
		count = (count + backtrack(c, 1)) % p
	}

	return count
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
