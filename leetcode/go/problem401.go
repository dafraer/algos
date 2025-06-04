	func readBinaryWatch(turnedOn int) []string {
		ans := make([]string, 0)
		if turnedOn > 8 {
			return ans
		}
		hours := []int{1, 2, 4, 8}
		mins := []int{32, 16, 8, 4, 2, 1}
		var dfs func(left int, h int, m int, ind int)
		dfs = func(left int, h int, m int, ind int) {
			if m > 59 || h > 11 {
				return
			}
			if left == 0 {
				ans = append(ans, format(h, m))
				return
			}
			if ind >= len(hours)+len(mins) {
				return
			}
			if ind >= len(hours) {
				dfs(left, h, m, ind+1)
				dfs(left-1, h, m+mins[ind-len(hours)], ind+1)
			} else {
				dfs(left, h, m, ind+1)
				dfs(left-1, h+hours[ind], m, ind+1)
			}
		}
		dfs(turnedOn, 0, 0, 0)
		return ans
	}

	func format(h, m int) string {
		var b strings.Builder
		b.WriteString(strconv.Itoa(h))
		if m < 10 {
			b.WriteString(":0")
		} else {
			b.WriteString(":")
		}
		b.WriteString(strconv.Itoa(m))
		return b.String()
	}
