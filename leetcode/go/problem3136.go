func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}
	cntVow := 0
	cntCons := 0
	b := []byte(word)
	for i := 0; i < len(b); i++ {
		if (b[i] < '0' || b[i] > '9') && !isLetter(b[i]) {
			return false
		}
		if isLetter(b[i]) && !isVowel(b[i]) {
			cntCons++
		} else if isVowel(b[i]) {
			cntVow++
		}
	}
	return cntVow > 0 && cntCons > 0
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'A' || b == 'E' || b == 'e' || b == 'i' || b == 'I' || b == 'o' || b == 'O' || b == 'u' || b == 'U'
}

func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}
