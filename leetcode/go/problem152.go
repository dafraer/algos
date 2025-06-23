	func maxProduct(nums []int) int {
		mx := nums[0]
		dp := make([][2]int, len(nums))
		dp[0][0] = nums[0]
		dp[0][1] = nums[0]
		for i := 1; i < len(dp); i++ {
			dp[i][0] = max(nums[i], max(dp[i-1][0]*nums[i], dp[i-1][1]*nums[i]))
			dp[i][1] = min(nums[i], min(dp[i-1][0]*nums[i], dp[i-1][1]*nums[i]))
			mx = max(mx, dp[i][0])
		}
		return mx
	}
