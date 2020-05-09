package main

import (
	"math"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	common := strs[0]
	// iterate words
	for i := 1; i < len(strs); i++ {
		lencommon := len(common)
		compare := strs[i]
		lencompare := len(compare)
		l := int(math.Min(float64(lencommon), float64(lencompare)))

		// iterate letters
		tempcommon := make([]byte, 0)
		for j := 0; j < l; j++ {
			letterCompare := compare[j]
			letterCommon := common[j]
			if letterCommon == letterCompare {
				tempcommon = append(tempcommon, letterCompare)
			} else {
				break
			}
		}
		common = string(tempcommon)
	}
	return common
}

func lcpDC(strs []string) string {
	l := len(strs)

	if l == 1 {
		return strs[0]
	}

	if l == 0 {
		return ""
	}
	middle := l / 2
	left := strs[:middle]
	right := strs[middle:]

	return compare(lcpDC(left), lcpDC(right))
}

func compare(left, right string) string {
	ll := len(left)
	lr := len(right)
	lmin := int(math.Min(float64(ll), float64(lr)))
	common := []byte{}

	for i := 0; i < lmin; i++ {
		if left[i] != right[i] {
			break
		} else {
			common = append(common, left[i])
		}
	}
	return string(common)
}

func lcpBS(strs []string) string {
	// edge cases
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	// determine min len
	lmin := 10000
	for _, v := range strs {
		lmin = int(math.Min(float64(lmin), float64(len(v))))
	}

	// divide and conquer
	// find the middle
	middle := lmin / 2 // this is where it starts dividing
	low := 0
	high := lmin
	// while middle is not repeated check if all strings in the array have that prefix
	for low <= high {
		tempmiddle := (low + high) / 2
		if isCommonPrefix(strs, tempmiddle) {
			// if all have then divide again btw middle and high
			// also we need to restrict the bounds
			low = tempmiddle + 1
			middle = tempmiddle
		} else {
			// if all don't have then divide again btw low and middle
			// also we need to restrict the bounds
			high = tempmiddle - 1
		}

	}

	return string(strs[0][:middle])
}

func isCommonPrefix(strs []string, middle int) bool {
	commonPrefix := string(strs[0][:middle])
	for _, v := range strs {
		if v[:middle] != commonPrefix {
			return false
		}
	}
	return true
}
