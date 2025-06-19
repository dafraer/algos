	func myPow(x float64, n int) float64 {
		if n >= 0 {
			return posPov(x, n)
		}
		return 1 / myPow(x, int(math.Abs(float64(n))))
	}

	func posPov(x float64, n int) float64 {
		if n == 0 {
			return 1
		}
		if n%2 == 0 {
			tmp := myPow(x, n/2)
			return tmp * tmp
		}
		return myPow(x, n-1) * x
	}
