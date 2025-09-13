	func largestGoodInteger(num string) string {
		mx := -1
		ans := ""
		for i := 0; i < len(num)-2; i++ {
			if num[i] == num[i+1] && num[i+1] == num[i+2] {
				if n, _ := strconv.Atoi(num[i : i+3]); n > mx {
					mx = n
					ans = num[i : i+3]
				}
			}
		}
		return ans
	}
