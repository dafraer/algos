	func containsNearbyDuplicate(nums []int, k int) bool {
		m := make(map[int]int)
		for i, v := range nums {
			if j, ok := m[v]; ok && math.Abs(float64(j-i)) <= float64(k) {
				return true
			}
			m[v] = i
		}
		return false
	}
