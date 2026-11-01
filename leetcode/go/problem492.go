func constructRectangle(area int) []int {
	mn := math.MaxInt
	l := 1
	for i := 2; area/i >= 2; i++ {
		if (area%i == 0) && (mn > int(math.Abs(float64(area/i-i)))) {
			mn = int(math.Abs(float64(area/i - i)))
			l = i
		}
	}
	return []int{area / l, l}
}