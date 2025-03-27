func getSneakyNumbers(nums []int) []int {
	ans := make([]int, 0, 2)
	m := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			ans = append(ans, nums[i])
		}
		m[nums[i]] = struct{}{}
	}
	return ans
}
