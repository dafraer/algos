func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		m := (r + l) / 2
		switch target {
		case nums[l]:
			return l
		case nums[m]:
			return m
		case nums[r]:
			return r
		}
		if nums[m] > nums[r] {
			if target < nums[r] || target > nums[m] {
				l = m + 1
			} else {
				r = m
			}

		} else {
			if target > nums[m] && target <= nums[r] {
				l = m + 1
			} else {
				r = m
			}
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}
