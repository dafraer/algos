func detectCapitalUse(word string) bool {
	if len(word) < 2 {
		return true
	}
	firstCapital := (word[0] >= 65 && word[0] <= 90) && !(word[1] >= 65 && word[1] <= 90)
	for i := 1; i < len(word); i++ {
		if (firstCapital && !(word[i] >= 97 && word[i] <= 122)) || (!firstCapital && (word[i] >= 97 && word[i] <= 122) != (word[i-1] >= 97 && word[i-1] <= 122)) {
			return false
		}
	}
	return true
}