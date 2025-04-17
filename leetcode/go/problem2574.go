func leftRightDifference(nums []int) []int {
	if len(nums) == 1 {
		return []int{0}
	}
	leftSum := make([]int, len(nums))
	rightSum := make([]int, len(nums))
	leftSum[1] = nums[0]
	for i := 2; i < len(nums); i++ {
		leftSum[i] = leftSum[i-1] + nums[i-1]
	}
	rightSum[len(nums)-2] = nums[len(nums)-1]
	for i := len(nums) - 3; i >= 0; i-- {
		rightSum[i] = rightSum[i+1] + nums[i+1]
	}
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = int(math.Abs(float64(leftSum[i] - rightSum[i])))
	}
	return ans
}
