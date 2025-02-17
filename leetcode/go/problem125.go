func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r {
		if !isAlphaNumerical(s[l]) {
			l++
		} else if !isAlphaNumerical(s[r]) {
			r--
		} else if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		} else {
			r--
			l++
		}
	}
	return true
}

func isAlphaNumerical(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}
