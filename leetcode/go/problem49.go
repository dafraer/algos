func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, v := range strs {
		var alp [26]int
		for _, s := range []byte(v) {
			alp[s-'a']++
		}
		m[alp] = append(m[alp], v)
	}
	ans := make([][]string, 0)
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}