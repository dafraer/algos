	func findDisappearedNumbers(nums []int) []int {
		a := make([]bool, len(nums)+1)
		for i := 0; i < len(nums); i++ {
			a[nums[i]] = true
		}
		ans := make([]int, 0, len(nums))
		for i := 1; i < len(a); i++ {
			if !a[i] {
				ans = append(ans, i)
			}
		}
		return ans
	}
