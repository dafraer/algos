func lengthLongestPath(input string) int {
	prefix := []int{0}
	longest := 0

	for _, line := range strings.Split(input, "\n") {
		name := strings.TrimLeft(line, "\t")
		depth := len(line) - len(name)
		length := prefix[depth] + len(name)

		if strings.Contains(name, ".") {
			longest = max(longest, length)
		} else {
			prefix = append(prefix[:depth+1], length+1)
		}
	}
	return longest
}

func split(s string, indent int) []string {
	sep := ""
	for _ = range indent {
		sep += "\t"
	}
	res := make([]string, 0)
	prevInd := 0
	for i := 0; i < len(s)-len(sep); i = max(i+1, prevInd) {
		if s[i:i+indent] == sep && s[i+indent+1] != '\n' && s[i+indent+1] != '\t' {
			res = append(res, s[prevInd:i])
			prevInd = i + indent + 1
		}
	}
	return res
}