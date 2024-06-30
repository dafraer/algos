func maxOperations(nums []int, k int) int {
	m := make(map[int]int)
	cnt := 0
	for _, v := range nums {
		m[v]++
	}
	for _, v := range nums {
		if k-v == v {
			if m[v] > 1 {
				cnt++
				m[k-v]--
				m[v]--
			}
		} else {
			if m[k-v] > 0 && m[v] > 0 {
				cnt++
				m[k-v]--
				m[v]--
			}
		}
	}
	return cnt
}
