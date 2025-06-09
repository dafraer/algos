	func arrangeCoins(n int) int {
		l, r := 0, n
		for l < r {
			m := (l + r) / 2
			if m*(m+1)/2 == n {
				return m
			} else if m*(m+1)/2 > n {
				r = m - 1
			} else {
				l = m + 1
			}
		}
		if r*(r+1)/2 > n {
			r--
		}
		return r
	}
