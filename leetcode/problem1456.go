func maxVowels(s string, k int) int {
	sum := 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			sum++
		}
	}
	mx := sum
	for i := k; i < len(s); i++ {
		if isVowel(s[i]) {
			sum++
		}
		if isVowel(s[i-k]) {
			sum--
		}
		mx = max(mx, sum)
	}
	return mx
}

func isVowel(b byte) bool {
	switch b {
	case 'a':
		return true
	case 'e':
		return true
	case 'i':
		return true
	case 'o':
		return true
	case 'u':
		return true
	default:
		return false
	}
}
