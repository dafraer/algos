	func largestLocal(grid [][]int) [][]int {
		ans := make([][]int, len(grid)-2)
		for i := 1; i < len(grid)-1; i++ {
			for j := 1; j < len(grid)-1; j++ {
				mx := 0
				for k := i - 1; k < (i-1)+3; k++ {
					for l := j - 1; l < (j-1)+3; l++ {
						mx = max(mx, grid[k][l])
					}
				}
				ans[i-1] = append(ans[i-1], mx)
			}
		}
		return ans
	}   
