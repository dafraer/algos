func maxProfit(prices []int) int {
	mx := prices[0]
	mn := prices[0]
	profit := 0
	for _, v := range prices {
		if v < mn {
			profit = max(profit, max(0, mx-mn))
			mn = v
			mx = 0
		}
		mx = max(mx, v)
	}
	return max(profit, mx-mn)
}