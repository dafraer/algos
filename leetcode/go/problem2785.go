	type letter struct {
		l   byte
		occ int
	}

	func sortVowels(s string) string {
		m := make(map[byte]int)
		for i := 0; i < len(s); i++ {
			if vowel(s[i]) {
				m[s[i]]++
			}
		}
		a := make([]letter, len(m))
		i := 0
		for k, v := range m {
			a[i] = letter{k, v}
			i++
		}
		sort.Slice(a, func(i, j int) bool { return a[i].l < a[j].l })
		curr := 0
		b := []byte(s)
		for i := 0; i < len(b); i++ {
			if vowel(s[i]) {
				b[i] = a[curr].l
				a[curr].occ--
				if a[curr].occ <= 0 {
					curr++
				}
			}
		}
		return string(b)
	}
	func vowel(b byte) bool {
		return b == 'a' || b == 'i' || b == 'e' || b == 'o' || b == 'u' || b == 'A' || b == 'I' || b == 'O' || b == 'E' || b == 'U'
	}
