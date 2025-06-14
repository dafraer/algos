	func letterCombinations(digits string) []string {
		if len(digits) == 0 {
			return []string{}
		}
		s := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
		ans := make([]string, 0)
		b := make([]byte, len(digits))
		var dfs func(i int)
		dfs = func(i int) {
			if i >= len(digits) {
				ans = append(ans, string(b))
				return
			}
			tmp, _ := strconv.Atoi(string(digits[i]))
			for j := 0; j < len(s[tmp]); j++ {
				b[i] = s[tmp][j]
				dfs(i + 1)
			}
		}
		dfs(0)
		return ans
	}
