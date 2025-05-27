	func combinationSum(candidates []int, target int) [][]int {
		return dfs(&candidates, target, []int{}, 0, 0)
	}

	func dfs(candidates *[]int, target int, comb []int, sum int, ind int) [][]int {
		if sum == target {
			return [][]int{comb}
		}
		if sum > target || ind >= len(*candidates) {
			return [][]int{}
		}
		combx := make([]int, len(comb))
		copy(combx, comb)
		newCombNoAdd := dfs(candidates, target, combx, sum, ind+1)
		comby := make([]int, len(comb)+1)
		copy(comby, comb)
		comby[len(comb)] = (*candidates)[ind]
		newCombAdd := dfs(candidates, target, comby, sum+(*candidates)[ind], ind)
		ans := make([][]int, 0)
		if len(newCombNoAdd) > 0 {
			ans = append(ans, newCombNoAdd...)
		}
		if len(newCombAdd) > 0 {
			ans = append(ans, newCombAdd...)
		}
		return ans
	}
