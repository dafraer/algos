func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	maxVolume := calcVolume(height[l], height[r], l, r)
	for r-l > 0 {
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
		maxVolume = max(maxVolume, calcVolume(height[l], height[r], l, r))
	}
	fmt.Println(l, r)
	return maxVolume
}

func calcVolume(h1, h2, l, r int) int {
	return min(h1, h2) * (r - l)
}
