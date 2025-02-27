func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	total := len(nums1) + len(nums2)
	half := total / 2
	l := -1
	r := len(nums1) - 1
	for {
		m := (r + l) / 2
		m2 := half - m - 2
		nums1L := math.MinInt
		nums1R := math.MaxInt
		nums2L := math.MinInt
		nums2R := math.MaxInt
		if m+1 < len(nums1) {
			nums1R = nums1[m+1]
		}

		if m >= 0 {
			nums1L = nums1[m]
		}

		if m2+1 < len(nums2) {
			nums2R = nums2[m2+1]
		}

		if m2 >= 0 {
			nums2L = nums2[m2]
		}

		if nums1L <= nums2R && nums2L <= nums1R {
			if total%2 == 0 {
				return float64(min(nums1R, nums2R)+max(nums1L, nums2L)) / 2
			} else {
				return float64(min(nums1R, nums2R))
			}
		} else if nums1L > nums2R {
			r = m - 1
		} else {
			l = m + 1
		}
	}
}  
