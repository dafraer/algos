	func isPowerOfTwo(n int) bool {
		b := 1
		for b < n {
			b = b << 1
		}
		return b == n
	}
