	func lengthOfLIS(nums []int) int {
		dp := make([]int, len(nums))
		dp[0] = 1
		mx := 1
		for i := 1; i < len(dp); i++ {
			dp[i] = 1
			for j := i - 1; j >= 0; j-- {
				if nums[j] < nums[i] {
					dp[i] = max(dp[j]+1, dp[i])
				}
			}
			mx = max(mx, dp[i])
		}
		return mx
	}
