	func permute(nums []int) [][]int {
		ans := make([][]int, 0, 620)
		var dfs func(comb []int)
		dfs = func(comb []int) {
			if len(comb) == len(nums) {
				ans = append(ans, comb)
				return
			}
			for i := 0; i < len(nums); i++ {
				if !inComb(comb, nums[i]) {
					tmp := make([]int, len(comb)+1)
					copy(tmp, comb)
					tmp[len(comb)] = nums[i]
					dfs(tmp)
				}
			}
		}
		dfs([]int{})
		return ans
	}

	func inComb(s []int, v int) bool {
		for i := 0; i < len(s); i++ {
			if s[i] == v {
				return true
			}
		}
		return false
	}
