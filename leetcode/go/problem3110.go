func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}
func scoreOfString(s string) int {
	ans := 0
	for i := 0; i < len(s)-1; i++ {
		fmt.Println(s[i])
		ans += abs(int(s[i]) - int(s[i+1]))
	}
	return ans
}
