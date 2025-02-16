func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})
	for _, v := range nums {
		m[v] = struct{}{}
	}
	mx := 0
	for k, _ := range m {
		if _, ok := m[k-1]; !ok {
			cnt := 1
			cur := k + 1
			for {
				if _, ok := m[cur]; ok {
					cnt++
					cur++
				} else {
					break
				}
			}
			mx = max(mx, cnt)
		}
	}
	return mx
}
