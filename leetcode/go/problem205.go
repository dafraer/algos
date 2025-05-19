	func isIsomorphic(s string, t string) bool {
		if len(s) != len(t) {
			return false
		}
		m := make(map[byte]byte)
		m1 := make(map[byte]byte)
		for i := 0; i < len(s); i++ {
			m[s[i]] = t[i]
			m1[t[i]] = s[i]
		}
		for i := 0; i < len(s); i++ {
			if m[s[i]] != t[i] || m1[t[i]] != s[i] {
				return false
			}
			m[s[i]] = t[i]
			m1[t[i]] = s[i]
		}
		return true
	}
