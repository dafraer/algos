func lengthOfLongestSubstring(s string) int {
	mx := 0
	l := 0
	m := make(map[byte]struct{})
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			for l < i {
				delete(m, s[l])
				if s[l] == s[i] {
					l++
					break
				}
				l++
			}
		}
		m[s[i]] = struct{}{}
		mx = max(mx, i-l+1)
	}
	return mx
}
