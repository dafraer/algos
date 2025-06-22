	func coinChange(coins []int, amount int) int {
		if amount == 0 {
			return 0
		}
		s := make([]int, amount+1)
		for i := 1; i < len(s); i++ {
			for j := 0; j < len(coins); j++ {
				if i-coins[j] == 0 {
					s[i] = 1
					break
				} else if i-coins[j] >= 1 {
					if s[i-coins[j]] != 0 && s[i] == 0 {
						s[i] = s[i-coins[j]] + 1
					} else if s[i-coins[j]] != 0 && s[i] != 0 {
						s[i] = min(s[i], s[i-coins[j]]+1)
					}
				}
			}
		}
		if s[amount] == 0 {
			return -1
		}
		return s[amount]
	}
