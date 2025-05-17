	func isHappy(n int) bool {
		s := strconv.Itoa(n)
		m := make(map[int]struct{})
		for n > 0 {
			sum := 0
			for i := 0; i < len(s); i++ {
				tmp, _ := strconv.Atoi(string(s[i]))
				sum += tmp * tmp
			}
			if _, ok := m[sum]; ok {
				return sum == 1
			}
			m[sum] = struct{}{}
			s = strconv.Itoa(sum)
		}
		return false
	}
