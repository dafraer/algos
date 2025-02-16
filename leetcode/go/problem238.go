func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	cnt, cntz, cntnoz := 1, 0, 1
	for i := 0; i < len(nums); i++ {
		cnt *= nums[i]
		if nums[i] != 0 {
			cntnoz *= nums[i]
		} else {
			cntz++
		}
	}
	if cntz > 1 {
		return ans
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			ans[i] = cntnoz
		} else {
			ans[i] = cnt / nums[i]
		}
	}
	return ans
}
