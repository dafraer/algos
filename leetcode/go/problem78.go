	func subsets(nums []int) [][]int {
		return dfs([]int{}, &nums, 0)
	}

	func dfs(set []int, nums *[]int, i int) [][]int {
		if i == len(*nums) {
			return [][]int{set}
		}
		ans := dfs(set, nums, i+1)
		newSet := make([]int, len(set))
		copy(newSet, set)
		newSet = append(newSet, (*nums)[i])
		ans = append(ans, dfs(newSet, nums, i+1)...)
		return ans
	}
