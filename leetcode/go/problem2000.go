	func reversePrefix(word string, ch byte) string {
		i := 0
		newS := ""
		for ; i < len(word); i++ {
			if word[i] == ch {
				if i < len(word)-1 {
					newS = word[i+1:]
				}
				break
			}
		}
		if i != len(word)-1 && newS == "" {
			return word
		}
		var b strings.Builder
		for j := i; j >= 0; j-- {
			b.WriteByte(word[j])
		}
		b.WriteString(newS)
		return b.String()
	}
