	func judgeCircle(moves string) bool {
		var pos [2]int
		for i := 0; i < len(moves); i++ {
			switch moves[i] {
			case 'U':
				pos[1]++
			case 'D':
				pos[1]--
			case 'L':
				pos[0]--
			case 'R':
				pos[0]++
			}
		}
		return pos[0] == 0 && pos[1] == 0
	}
