func equalPairs(grid [][]int) int {
    m := make(map[string]int)
    cnt := 0
    for i:=0; i < len(grid); i++ {
        m[str(grid[i])]++
    }

    for i:=0; i < len(grid); i++ {
        arr := make([]int, len(grid))
        for j:=0; j < len(grid);j++ {
            arr[j] = grid[j][i]
        }
            cnt+=m[str(arr)]
    }
    return cnt
}

func str(s []int) string {
	return fmt.Sprintf("%v", s)
}
