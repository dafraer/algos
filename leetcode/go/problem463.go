	func islandPerimeter(grid [][]int) int {
		ans := 0
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] == 0 {
					continue
				}
				if i+1 >= len(grid) || grid[i+1][j] == 0 {
					ans++
				}
				if i-1 < 0 || grid[i-1][j] == 0 {
					ans++
				}
				if j+1 >= len(grid[i]) || grid[i][j+1] == 0 {
					ans++
				}
				if j-1 < 0 || grid[i][j-1] == 0 {
					ans++
				}
			}
		}
		return ans
	}
