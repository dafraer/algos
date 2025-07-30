	func longestSubarray(nums []int) int {
		mx := 0
		for i := 0; i < len(nums); i++ {
			mx = max(nums[i], mx)
		}
		cnt := 0
		mxCnt := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] == mx {
				cnt++
			} else {
				mxCnt = max(mxCnt, cnt)
				cnt = 0
			}
		}
		return max(mxCnt, cnt)
	}
