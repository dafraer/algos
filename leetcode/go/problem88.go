func merge(nums1 []int, m int, nums2 []int, n int) {
	s := make([]int, m)
	copy(s, nums1[:m])
	i := 0
	j := 0
	k := 0
	for i < m && j < n {
		if s[i] < nums2[j] {
			nums1[k] = s[i]
			i++
		} else {
			nums1[k] = nums[j]
			j++
		}
		k++
	}
	for i < m {
		nums1[k] = s[i]
		i++
		k++
	}
	for j < n {
		nums1[k] = nums2[j]
		j++
		k++
	}
}
