func threeSum(nums []int) [][]int {
	ans := make([][]int, 0, len(nums))
	sort.Ints(nums)
	for i := 0; i < len(nums) && nums[i] < 1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		cur := nums[i]
		des := -cur
		l := i + 1
		r := len(nums) - 1
		used1, used2 := math.MinInt, math.MinInt
		for l < r {
			if nums[l]+nums[r] < des {
				l++
			} else if nums[l]+nums[r] > des {
				r--
			} else {
				if !(nums[l] == used1 && nums[r] == used2) {
					ans = append(ans, []int{cur, nums[l], nums[r]})
					used1 = nums[l]
					used2 = nums[r]
				}
				r--
			}
		}

	}
	return ans
}
