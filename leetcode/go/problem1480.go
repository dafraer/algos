func runningSum(nums []int) []int {
	rsum := make([]int, len(nums))
	rsum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		rsum[i] = rsum[i-1] + nums[i]
	}
	return rsum
}
