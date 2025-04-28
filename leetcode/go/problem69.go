func mySqrt(x int) int {
	l := 0
	r := x
	for l < r {
		m := (r + l) / 2
		if m*m == x {
			return m
		} else if m*m > x {
			r = m
		} else {
			l = m + 1
		}
	}
	if l*l > x {
		return l - 1
	}
	return l
}
