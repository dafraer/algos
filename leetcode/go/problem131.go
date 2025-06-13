	func isPalindrome(s string) bool {
		l := 0
		r := len(s) - 1
		for l < r {
			if !isAlphaNumerical(s[l]) {
				l++
			} else if !isAlphaNumerical(s[r]) {
				r--
			} else if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
				return false
			} else {
				r--
				l++
			}
		}
		return true
	}
	func isAlphaNumerical(b byte) bool {
		return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
	}
	func partition(s string) [][]string {
		var dfs func(i int)
		ans := make([][]string, 0)
		part := make([]string, 0)
		dfs = func(i int) {
			if i >= len(s) {
				newS := make([]string, len(part))
				copy(newS, part)
				ans = append(ans, newS)
				return
			}
			for j := i; j < len(s); j++ {
				b := []byte(s)
				newS := string(b[i : j+1])
				if isPalindrome(newS) {
					part = append(part, newS)
					dfs(j + 1)
					part = part[:len(part)-1]
				}
			}
		}
		dfs(0)
		return ans
	}
