func minEatingSpeed(piles []int, h int) int {
	l := 1
	r := 0
	for _, v := range piles {
		r = max(r, v)
	}
	r++
	for l < r {
		m := (r + l) / 2
		cnt := 0
		for _, v := range piles {
			cnt += int(math.Ceil(float64(v) / float64(m)))
		}
		if cnt <= h {
			r = m
		} else if cnt > h {
			l = m + 1
		}
	}
	return l

}
