
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	a := make([][200]int, len(triangle))
	ans := math.MaxInt
	a[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j > 0 && j < len(triangle[i])-1 {
				a[i][j] = min(a[i-1][j-1], a[i-1][j]) + triangle[i][j]
			} else if j > 0 {
				a[i][j] = a[i-1][j-1] + triangle[i][j]
			} else {
				a[i][j] = a[i-1][j] + triangle[i][j]
			}
			if i == len(triangle)-1 {
				ans = min(ans, a[i][j])
			}
		}
	}
	return ans
}
