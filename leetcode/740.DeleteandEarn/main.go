package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func deleteAndEarn2(nums []int) int {
	var maxNr int
	m := make(map[int]int)
	memo := make(map[int]int)
	for _, v := range nums {
		if v > maxNr {
			maxNr = v
		}

		m[v] += v
	}

	var dp func(i int) int
	dp = func(i int) int {
		if v, ok := memo[i]; ok {
			return v
		}

		var res int
		if i == 0 {
			res = 0
		} else if i == 1 {
			res = m[i]
		} else {
			res = max(dp(i-2)+m[i], dp(i-1))
		}

		memo[i] = res

		return res
	}

	return dp(maxNr)
}

func deleteAndEarn(nums []int) int {
	var prevprev, prev, maxNr int
	m := make(map[int]int)
	for _, v := range nums {
		if v > maxNr {
			maxNr = v
		}

		m[v] += v
	}

	prevprev = 0
	prev = m[1]

	for i := 2; i <= maxNr; i++ {
		temp := prev
		prev = max(m[i]+prevprev, prev)
		prevprev = temp
	}

	return prev
}

func main() {
	nums := []int{3, 4, 2}
	fmt.Println(deleteAndEarn(nums))
	nums2 := []int{2, 2, 3, 3, 3, 4}
	fmt.Println(deleteAndEarn(nums2))
	nums3 := []int{5, 4, 5, 4, 3, 5, 3}
	fmt.Println(deleteAndEarn(nums3))
	nums4 := []int{5, 5, 5, 6, 6, 6, 6, 7, 7}
	fmt.Println(deleteAndEarn(nums4))

}
