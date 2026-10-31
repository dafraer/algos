func findMaxConsecutiveOnes(nums []int) int {
	mx := 0
	cur := 0
	for i := range len(nums) {
		if nums[i] == 0 {
			if cur > mx {
				mx = cur
			}
			cur = 0
		} else {
			cur += 1
		}
	}
	return max(cur, mx)
}