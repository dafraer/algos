	func numDecodings(s string) int {
		a := 1
		b := 0
		if s[0] != '0' {
			b = 1
		}
		for i := 2; i <= len(s); i++ {
			cur := 0
			d, _ := strconv.Atoi(s[i-1 : i])
			if d >= 1 && d <= 9 {
				cur += b
			}

			dd, _ := strconv.Atoi(s[i-2 : i])
			if dd >= 10 && dd <= 26 {
				cur += a
			}
			a, b = b, cur
		}
		return b
	}
