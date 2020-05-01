package factorial

func Rfactorial(n int) int {
	if n == 2 {
		return 2
	}

	return n * Rfactorial(n-1)
}
