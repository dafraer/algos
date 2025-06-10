	func rob(nums []int) int {
		if len(nums) == 1 {
			return nums[0]
		}
		s := make([]int, len(nums))
		s[0] = nums[0]
		s[1] = max(nums[1], nums[0])
		for i := 2; i < len(nums); i++ {
			s[i] = max(s[i-1], s[i-2]+nums[i])
		}
		return s[len(s)-1]
	}
