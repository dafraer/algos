	func isPowerOfThree(n int) bool {
		p := 1
		for p < n {
			p *= 3
		}
		return p == n
	}
