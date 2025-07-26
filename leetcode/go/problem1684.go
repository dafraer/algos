	func countConsistentStrings(allowed string, words []string) int {
		cnt := 0
		m := make(map[byte]struct{})
		for i := 0; i < len(allowed); i++ {
			m[allowed[i]] = struct{}{}
		}
	loop:
		for _, v := range words {
			for i := 0; i < len(v); i++ {
				if _, ok := m[v[i]]; !ok {
					continue loop
				}
			}
			cnt++
		}
		return cnt
	}
