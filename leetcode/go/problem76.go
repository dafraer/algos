func minWindow(s string, t string) string {
	tmap := make(map[byte]int)
	smap := make(map[byte]int)
	need := len(t)
	for i := 0; i < len(t); i++ {
		tmap[t[i]]++
	}
	l := 0
	mn := ""
	mnlen := math.MaxInt
	have := 0
	for i := 0; i < len(s); i++ {
		smap[s[i]]++
		if tmap[s[i]] > 0 && smap[s[i]] == tmap[s[i]] {
			have += tmap[s[i]]
		}
		for have == need {
			if mnlen > i-l {
				mnlen = i - l
				mn = s[l : i+1]
			}
			smap[s[l]]--
			if smap[s[l]] < tmap[s[l]] {
				have -= tmap[s[l]]
			}
			l++
		}
	}
	return mn
}
