package fibonacci

var cache = map[int]int{
	1: 1,
	2: 1,
}

// Rfibonacci recursive ribo
func Rfibonacci(n int) int {

	if cache[n] != 0 {
		return cache[n]
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 1
	}

	cache[n] = Rfibonacci(n-1) + Rfibonacci(n-2)
	return cache[n]
}
