	func exist(board [][]byte, word string) bool {
		var dfs func(i, j, ind int) bool
		visited := make([][]bool, len(board))
		for i, _ := range board {
			visited[i] = make([]bool, len(board[i]))
		}
		dfs = func(i, j, ind int) bool {
			if ind == len(word) {
				return true
			}
			if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
				return false
			}
			if board[i][j] != word[ind] || visited[i][j] {
				return false
			}
			visited[i][j] = true
			res := dfs(i-1, j, ind+1) || dfs(i+1, j, ind+1) || dfs(i, j+1, ind+1) || dfs(i, j-1, ind+1)
			if !res {
				visited[i][j] = false
				return false
			}
			return true
		}
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
		return false
	}
