func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	ans := make([][]int, 0, rowIndex+1)
	ans = append(ans, []int{1}, []int{1, 1})
	for i := 3; i <= rowIndex+1; i++ {
		tmp := make([]int, i)
		tmp[0] = 1
		tmp[i-1] = 1
		for j := 1; j < i-1; j++ {
			tmp[j] = ans[i-2][j-1] + ans[i-2][j]
		}
		ans = append(ans, tmp)
	}
	return ans[rowIndex]
}
