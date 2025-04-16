func maximumWealth(accounts [][]int) int {
	mx := 0
	for i := 0; i < len(accounts); i++ {
		cnt := 0
		for j := 0; j < len(accounts[i]); j++ {
			cnt += accounts[i][j]
		}
		mx = max(mx, cnt)
	}
	return mx
}
