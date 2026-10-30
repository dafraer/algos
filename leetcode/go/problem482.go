func licenseKeyFormatting(s string, k int) string {
	s = strings.ToUpper(s)
	b := make([]byte, 0, len(s))
	for i := range len(s) {
		if s[i] != '-' {
			b = append(b, s[i])
		}
	}
	if len(b) == 0 {
		return ""
	}
	first := len(b) % k

	res := make([]byte, 0, len(b)+len(b)/k-1)
	prev := 0
	for i := 0; i < len(b); i++ {
		if i != 0 && (i == first || i-prev == k) {
			res = append(res, '-')
			prev = i
		}
		res = append(res, b[i])
	}
	return string(res)
}		