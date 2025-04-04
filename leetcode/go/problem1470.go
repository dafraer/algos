func shuffle(nums []int, n int) []int {
	ans := make([]int, 2*n)
	cntx := 0
	cnty := n
	for i := 0; i < 2*n; i++ {
		if i%2 == 0 {
			ans[i] = nums[cntx]
			cntx++
		} else {
			ans[i] = nums[cnty]
			cnty++
		}
	}
	return ans
}
