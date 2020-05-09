package main

import "fmt"

func main() {
	fmt.Println("LEETCODE")

	// Two Sum
	fmt.Println("Two Sum:", twoSum([]int{2, 7, 11, 19}, 9))

	// Reverse Integer
	fmt.Println("Reverse Integer:", reverseInteger(-123))

	// Palindrome Number
	fmt.Println("Palindrome Number:", isPalindrome(121))

	// Roman to Integer
	fmt.Println("Roman to Integer:", romanToInt("III"))

	// Longest Common Prefix
	fmt.Println("Longest Common Prefix:", lcpBS([]string{"flower", "flow", "flight", "flood"}))

	// Valid Parentheses
	fmt.Println("Valid Parentheses:", isValid("()[]{}"))

	// Merge two sorted lists
	fmt.Println("Merge two sorted lists:", mergeTwoLists(&ListNode{1, &ListNode{2, &ListNode{4, nil}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}))
}
