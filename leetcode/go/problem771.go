func numJewelsInStones(jewels string, stones string) int {
	cnt := 0
	for _, v := range stones {
		if strings.Contains(jewels, string(v)) {
			cnt++
		}
	}
	return cnt
}
