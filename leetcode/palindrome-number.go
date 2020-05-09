package main

// IsPalindrome palindrome
func isPalindrome(x int) bool {
	if x < 0 || x%10 == 0 {
		return false
	}
	var res int = 0
	for x > res {
		pop := x % 10
		x /= 10
		res = res*10 + pop
	}

	return res == x || x == res/10
}
