func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		m := (r + l) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}
