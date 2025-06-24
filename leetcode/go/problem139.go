	func wordBreak(s string, wordDict []string) bool {
		dp := make([]bool, len(s)+1)
		dp[len(s)] = true
		for i := len(s) - 1; i >= 0; i-- {
			for _, v := range wordDict {
				if len(v)+i <= len(s) && s[i:i+len(v)] == v {
					dp[i] = dp[i+len(v)]
				}
				if dp[i] {
					break
				}
			}
		}
		return dp[0]
	}
