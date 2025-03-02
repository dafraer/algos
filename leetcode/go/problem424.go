func characterReplacement(s string, k int) int {
	m := make(map[byte]int)
	mxf := 0
	res := 0
	l := 0
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		mxf = max(mxf, m[s[i]])
		if (i-l+1)-mxf <= k {
			res = max(i-l+1, res)
		} else {
			m[s[l]]--
			l++
		}
	}
	return res
}
