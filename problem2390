func removeStars(s string) string {
	b := []byte(s)
	cnt := 0
	ans := make([]byte, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '*' {
			cnt++
			b[i] = '-'
		} else if cnt > 0 {
			cnt--
			b[i] = '-'
		}
	}

	for i := 0; i < len(b); i++ {
		if b[i] != '-' {
			ans = append(ans, b[i])
		}
	}
	return string(ans)
}
