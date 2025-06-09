	func intersection(nums1 []int, nums2 []int) []int {
		m := make(map[int]struct{})
		for _, v := range nums1 {
			m[v] = struct{}{}
		}
		ans := make([]int, 0, min(len(nums1), len(nums2)))
		for _, v := range nums2 {
			if _, ok := m[v]; ok {
				ans = append(ans, v)
				delete(m, v)
			}
		}
		return ans
	}
