func numIdenticalPairs(nums []int) int {
	m := make(map[int]int)
	cnt := 0
	for i := 0; i < len(nums); i++ {
		cnt += m[nums[i]]
		m[nums[i]]++
	}
	return cnt
}
