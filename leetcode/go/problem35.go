func searchInsert(nums []int, target int) int {
	r := len(nums)
	l := 0
	for l < r {
		m := (r + l) / 2
		switch {
		case nums[m] == target:
			return m
		case nums[m] > target:
			r = m
		case nums[m] < target:
			l = m + 1
		}
	}
	return l
}
