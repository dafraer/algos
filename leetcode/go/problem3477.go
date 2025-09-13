	func numOfUnplacedFruits(fruits []int, baskets []int) int {
		cnt := len(fruits)
		for i := 0; i < len(fruits); i++ {
			for j := 0; j < len(baskets); j++ {
				if baskets[j] >= fruits[i] {
					baskets[j] = -1
					cnt--
					break
				}
			}
		}
		return cnt
	}
