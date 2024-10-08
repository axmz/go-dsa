package main

import "fmt"

var count int

func combine1(n int, k int) [][]int {
	var res [][]int

	var backtrack func(x int, combination []int)
	backtrack = func(x int, combination []int) {
		count++
		fmt.Println(count)
		if len(combination) == k {
			res = append(res, combination)
			return
		}

		if x > n {
			return
		}

		for i := x; i <= n; i++ {
			if len(combination)+(n-i)+1 < k {
				break
			}
			c := make([]int, len(combination))
			copy(c, combination)
			c = append(c, i)
			backtrack(i+1, c)
		}

	}

	backtrack(1, []int{})
	return res
}

func combine(n int, k int) [][]int {
	var ans [][]int
	var curr []int
	var backtrack func(int)
	backtrack = func(firstNum int) {
		count++
		fmt.Println(count)
		if len(curr) == k {
			ans = append(ans, append([]int{}, curr...))
			return
		}
		need := k - len(curr)
		remain := n - firstNum + 1
		available := remain - need
		for num := firstNum; num <= firstNum+available; num++ {
			curr = append(curr, num)
			backtrack(num + 1)
			curr = curr[:len(curr)-1]
		}
	}
	backtrack(1)
	return ans
}

func main() {
	n := 6
	k := 2
	fmt.Println(combine1(n, k))
}
