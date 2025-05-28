func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var dfs func(comb []int, sum int, ind int)
	ans := make([][]int, 0)
	dfs = func(comb []int, sum int, ind int) {
		if sum == target {
			ans = append(ans, comb)
			return
		}
		if sum > target || ind >= len(candidates) {
			return
		}
		i := ind
		for i < len(candidates) && candidates[ind] == candidates[i] {
			i++
		}
		if sum+candidates[ind] <= target {
			tmp := make([]int, len(comb)+1)
			copy(tmp, comb)
			tmp[len(comb)] = candidates[ind]
			dfs(tmp, sum+candidates[ind], ind+1)
		}
		dfs(comb, sum, i)
	}
	dfs([]int{}, 0, 0)
	return ans
}



