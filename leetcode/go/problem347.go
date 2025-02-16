func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	mx := math.MinInt
	for _, v := range nums {
		m[v]++
		mx = max(mx, m[v])
	}
	cnt := make([][]int, mx+1)
	for k, v := range m {
		cnt[v] = append(cnt[v], k)
	}
	ans := make([]int, k)
	j := 0
	for i := mx; i >= 0 && j < k; i-- {
		for _, v := range cnt[i] {
			ans[j] = v
			j++
		}
	}
	return ans
}
