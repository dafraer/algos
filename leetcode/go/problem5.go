	func longestPalindrome(s string) string {
		if len(s) == 1 {
			return s
		}
		if len(s) == 2 {
			if s[0] != s[1] {
				return string(s[0])
			}
			return s
		}
		mx := 1
		ansL, ansR := 0, 0
		b := []byte(s)
		for i := 0; i < len(b); i++ {
			l, r := i, i
			for l >= 0 && r < len(b) && s[l] == s[r] {
				if r-l+1 > mx {
					mx = r - l + 1
					ansL, ansR = l, r
				}
				l--
				r++
			}
			l, r = i, i+1
			for l >= 0 && r < len(b) && s[l] == s[r] {
				if r-l+1 > mx {
					mx = r - l + 1
					ansL, ansR = l, r
				}
				l--
				r++
			}
		}
		return string(b[ansL : ansR+1])
	}
