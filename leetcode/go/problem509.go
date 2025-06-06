	func fib(n int) int {
		if n == 0 {
			return 0
		}
		s := make([]int, n+1)
		s[0] = 0
		s[1] = 1
		for i := 2; i < len(s); i++ {
			s[i] = s[i-1] + s[i-2]
		}
		return s[n]
	}
