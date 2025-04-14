func minimumOperations(nums []int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		cnt += min(nums[i]%3, (nums[i]-nums[i]%3)+3-nums[i])
	}
	return cnt
}
