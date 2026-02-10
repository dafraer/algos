func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	ans := make([][]byte, numRows)
	for i := 0; i < len(s); {
		for j := 0; j < numRows && i < len(s); j++ {
			ans[j] = append(ans[j], s[i])
			i++
		}
		for j := numRows - 2; j > 0 && i < len(s); j-- {
			ans[j] = append(ans[j], s[i])
			i++
		}
	}
	b := make([]byte, 0, len(s))
	for _, v := range ans {
		b = append(b, v...)
	}
	return string(b)
}


