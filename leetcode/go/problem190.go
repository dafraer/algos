	func reverseBits(num uint32) uint32 {
		s := formatInt(num)
		var ans uint32
		fmt.Println(s)
		for i := 0; i < len(s); i++ {
			if s[i] {
				ans += (1 << i)
			}
		}
		return ans
	}

	func formatInt(num uint32) [32]bool {
		var s [32]bool
		for i := 0; num > 0; i++ {
			if num%2 == 1 {
				s[31-i] = true
			}
			num /= 2
		}
		return s
	}
