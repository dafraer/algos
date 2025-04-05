func transformArray(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			nums[i] = 0
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i] = 1
		}
	}
	sort.Ints(nums)
	return nums
}
