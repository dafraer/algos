func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0, 1000)
	var dfs func(set []int, i int)
	dfs = func(set []int, i int) {
		if i == len(nums) {
			ans = append(ans, set)
			return
		}
		res1 := make([]int, len(set))
		copy(res1, set)
		ind := i
		for ind < len(nums) && nums[ind] == nums[i] {
			ind++
		}
		dfs(set, ind)
		tmp := make([]int, len(set)+1)
		copy(tmp, set)
		tmp[len(set)] = nums[i]
		dfs(tmp, i+1)
	}
	dfs([]int{}, 0)
	return ans
}

