func searchMatrix(matrix [][]int, target int) bool {
	l := 0
	r := (len(matrix) * len(matrix[0]))
	for l < r {
		m := (l + r) / 2
		if target > matrix[m/len(matrix[0])][m%len(matrix[0])] {
			l = m + 1
		} else if target < matrix[m/len(matrix[0])][m%len(matrix[0])] {
			r = m
		} else {
			return true
		}
	}
	return false
}