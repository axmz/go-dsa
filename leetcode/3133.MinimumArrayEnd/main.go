package main

func minEnd(n int, x int) int64 {
	// Fill the zero bits of x with bits from (n-1), from low to high.
	// This builds the smallest possible last value in O(log n).
	res := int64(x)
	remaining := int64(n - 1)
	mask := int64(1)

	for remaining > 0 {
		if (int64(x) & mask) == 0 {
			if (remaining & 1) == 1 {
				res |= mask
			}
			remaining >>= 1
		}
		mask <<= 1
	}

	return res
}
