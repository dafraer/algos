	func thirdMax(nums []int) int {
		sort.Ints(nums)
		s := make([]int, 1, len(nums))
		s[0] = nums[0]
		for i := 1; i < len(nums); i++ {
			if nums[i] != nums[i-1] {
				s = append(s, nums[i])
			}
		}
		if len(s) < 3 {
			return s[len(s)-1]
		}
		return s[len(s)-3]
	}
