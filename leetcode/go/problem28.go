func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if needle[0] == haystack[i] {
			j := 1
			k := i + 1
			isEqual := true
			for j < len(needle) && k < len(haystack) {
				if needle[j] != haystack[k] {
					isEqual = false
				}
				k++
				j++
			}
			if isEqual {
				return i
			}
		}
	}
	return -1
}
