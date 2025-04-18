func smallerNumbersThanCurrent(nums []int) []int {
	m := make(map[int]int)
	ans := make([]int, len(nums))
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)
	for i := len(nums) - 1; i >= 0; i-- {
		m[sorted[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		ans[i] = m[nums[i]]
	}
	return ans
}
