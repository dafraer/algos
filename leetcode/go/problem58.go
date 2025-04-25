func lengthOfLastWord(s string) int {
	strs := make([]strings.Builder, 0)
	j := 0
	flag := false
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if flag {
				continue
			}
			flag = true
			j++
		} else {
			for len(strs) < j+1 {
				strs = append(strs, strings.Builder{})
			}
			strs[j].WriteByte(s[i])
			flag = false
		}
	}
	return len(strs[len(strs)-1].String())
}
