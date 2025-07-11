	func countSubstrings(s string) int {
		if len(s) == 1 {
			return 1
		}
		if len(s) == 2 {
			if s[0] != s[1] {
				return 2
			}
			return 3
		}
		ans := 0
		for i := 0; i < len(s); i++ {
			l, r := i, i
			for l >= 0 && r < len(s) && s[l] == s[r] {
				ans++
				l--
				r++
			}
			l, r = i, i+1
			for l >= 0 && r < len(s) && s[l] == s[r] {
				ans++
				l--
				r++
			}
		}
		return ans
	}
