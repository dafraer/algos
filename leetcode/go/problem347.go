func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	cnt := make([][]int, len(nums)+1)
	for k, v := range m {
		cnt[v] = append(cnt[v], k)
	}
	ans := make([]int, k)
	j := 0
	for i := len(nums); i > 0 && j < k; i-- {
		for _, v := range cnt[i] {
			ans[j] = v
			j++
		}
	}
	return ans
}
