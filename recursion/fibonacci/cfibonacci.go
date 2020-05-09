package fibonacci

// FiboGenerator creates a closures fibo
func FiboGenerator() func() int {
	f1 := 0
	f2 := 1

	return func() int {
		temp := f1 + f2
		f1 = f2
		f2 = temp
		return f2
	}
}

// Cfibonacci closures fibo
func Cfibonacci(n int) int {
	fibo := FiboGenerator()
	res := 0
	for i := 1; i < n; i++ {
		res = fibo()
	}

	return res
}
