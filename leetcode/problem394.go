func decodeString(s string) string {
	ans := make([]byte, 0, len(s))
	b := []byte(s)
	brackets := 0
	inBrackets := 0
	num := make([]byte, 0, 3)
	for i := 0; i < len(s); i++ {
		switch {
		case b[i] == '[':
			if inBrackets <= 0 {
				brackets = i + 1
			}
			inBrackets++
		case b[i] == ']':
			inBrackets--
			if inBrackets <= 0 {
				for j := 0; j < toInt(num); j++ {
					ans = append(ans, []byte(decodeString(string(b[brackets:i])))...)
				}
				num = make([]byte, 0, 3)
			}
		default:
			if b[i] <= '9' && b[i] >= '0' && inBrackets <= 0 {
				num = append(num, b[i])
			} else if inBrackets <= 0 {
				ans = append(ans, b[i])
			}

		}
	}
	return string(ans)
}

func toInt(b []byte) int {
	x, _ := strconv.Atoi(string(b))
	return x
}
