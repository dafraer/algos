func dp(open, closed, n int, s string) []string {
	ans := make([]string, 0)
	if open > closed && closed < n {
		ans = append(ans, dp(open, closed+1, n, s+")")...)
	}
	if open < n {
		ans = append(ans, dp(open+1, closed, n, s+"(")...)
	}
	if open == n && closed == n {
		ans = append(ans, s)
	}
	return ans
}

func generateParenthesis(n int) []string {
	return dp(0, 0, n, "")
}
