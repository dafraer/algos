	import (
		"fmt"
		"strconv"
	)

	func myAtoi(s string) int {
		nString := make([]byte, 0, len(s))
		start := false
		for i := 0; i < len(s); i++ {
			if !(number(s[i]) || s[i] == '-' || s[i] == '+' || s[i] == ' ') {
				break
			}
			if start && !(number(s[i])) {
				break
			}
			if number(s[i]) || s[i] == '-' || s[i] == '+' {
				start = true
			}
			if start {
				nString = append(nString, s[i])
			}
		}
		if len(nString) == 0 {
			return 0
		}
		realNString := make([]byte, 0, len(s))
		j := 0
		if nString[0] == '-' || nString[0] == '+' {
			realNString = append(realNString, nString[0])
			j++
		}
		for j < len(nString) && nString[j] == '0' {
			j++
		}
		realNString = append(realNString, nString[j:]...)
		nString = realNString
		mx := 2147483647
		mn := -2147483648
		fmt.Println(string(nString))
		n, _ := strconv.Atoi(string(nString))
		fmt.Println(n)
		switch {
		case len(nString) > 11 && nString[0] == '-':
			return mn
		case len(nString) > 11 && nString[0] == '+':
			return mx
		case len(nString) > 11 || n > mx:
			return mx
		case n < mn:
			return mn
		}
		return n
	}

	func number(b byte) bool {
		return (b >= '0' && b <= '9')
	}
