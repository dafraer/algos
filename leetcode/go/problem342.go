	func isPowerOfFour(n int) bool {
		for i := 0; n > 0; i++ {
			if n>>1 == 0 && n&1 == 1 && i%2 == 0 {
				return true
			}
			if n&1 == 1 {
				return false
			}
			n = n >> 1
		}
		return false
	}
