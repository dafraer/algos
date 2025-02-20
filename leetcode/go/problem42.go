func trap(height []int) int {
	l1, r1 := 0, len(height)-1
	mxhl, mxhr := 0, 0
	vol := 0
	oldmx := 0
	for l1 < r1 {
		if height[l1] > mxhl {
			mxhl = height[l1]
		}
		if height[r1] > mxhr {
			mxhr = height[r1]
		}
		if oldmx < min(mxhr, mxhl) {
			taken := 0
			for i := l1; i < r1; i++ {
				if height[i]-oldmx > 0 {
					taken += min(min(mxhr, mxhl), height[i]) - oldmx
				}
			}
			vol += min(mxhr, mxhl)*(r1-l1) - taken - (r1-l1)*oldmx
		}
		if mxhl < mxhr {
			l1++
		} else {
			r1--
		}
		oldmx = min(mxhl, mxhr)
	}
	return vol
}
