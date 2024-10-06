func compressedString(word string) string {
	if len(word) == 0 {
		return ""
	}

	var b strings.Builder
	cnt := 1

	for i := 1; i <= len(word); i++ {
		switch {
		case i == len(word):
			b.WriteString(strconv.Itoa(cnt))
			b.WriteByte(word[i-1])
		case word[i] != word[i-1]:
			b.WriteString(strconv.Itoa(cnt))
			b.WriteByte(word[i-1])
			cnt = 1
		case cnt == 9:
			b.WriteString(strconv.Itoa(cnt))
			b.WriteByte(word[i-1])
			cnt = 1
		default:
			cnt++
		}
	}
	return b.String()
}

