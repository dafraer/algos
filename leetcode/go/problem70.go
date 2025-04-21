func climbStairs(n int) int {
	s := make([]int, n+1)
	return t(n, &s)
}

func t(n int, res *[]int) int {
	if n < 2 {
		return 1
	}
	if (*res)[n-1] == 0 {
		(*res)[n-1] = t(n-1, res)
	}
	if (*res)[n-2] == 0 {
		(*res)[n-2] = t(n-2, res)
	}
	(*res)[n] = (*res)[n-1] + (*res)[n-2]
	return (*res)[n]
}
