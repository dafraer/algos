	func maximumLength(nums []int) int {
		if len(nums) == 2 {
			return 2
		}
		dp := make([]int, len(nums))
		dp[0] = 1
		for i := 1; i < len(nums); i++ {
			dp[i] = dp[i-1]
			if nums[i-1]%2 != nums[i]%2 {
				dp[i]++
			}
		}
		cntOdd := 0
		cntEven := 0
		for i := 0; i < len(nums); i++ {
			if nums[i]%2 == 0 {
				cntEven++
			} else {
				cntOdd++
			}
		}
		return max(cntOdd, max(cntEven, dp[len(dp)-1]))
	}
