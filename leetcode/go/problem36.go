func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		m := make(map[byte]struct{})
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if _, ok := m[board[i][j]]; ok {
					return false
				}
				m[board[i][j]] = struct{}{}
			}
		}
	}
	for j := 0; j < 9; j++ {
		m := make(map[byte]struct{})
		for i := 0; i < 9; i++ {
			if board[i][j] != '.' {
				if _, ok := m[board[i][j]]; ok {
					return false
				}
				m[board[i][j]] = struct{}{}
			}
		}
	}
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			m := make(map[byte]struct{})
			for k := i; k < i+3; k++ {
				for l := j; l < j+3; l++ {
					if board[k][l] != '.' {
						if _, ok := m[board[k][l]]; ok {
							return false
						}
						m[board[k][l]] = struct{}{}
					}
				}
			}
		}
	}
	return true
}
