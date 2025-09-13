	func nextGreaterElement(nums1 []int, nums2 []int) []int {
		m := make(map[int]int)
		for i := 0; i < len(nums2); i++ {
			m[nums2[i]] = i
		}
		ans := make([]int, len(nums1))
		for i := 0; i < len(ans); i++ {
			ans[i] = -1
		}
		for i := 0; i < len(nums1); i++ {
			for j := m[nums1[i]] + 1; j < len(nums2); j++ {
				if nums2[j] > nums1[i] {
					ans[i] = nums2[j]
					break
				}
			}
		}
		return ans
	}
