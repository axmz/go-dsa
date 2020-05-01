package factorial

func Ifactorial(n int) int {
	var result int = 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	return result
}
