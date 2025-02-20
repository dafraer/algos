func trap(height []int) int {
	l, r := 0, len(height)-1
	mxhl, mxhr := 0, 0
	vol := 0
	for l < r {
		//Adjust max height variables
		mxhl = max(height[l], mxhl)
		mxhr = max(height[r], mxhr)

		//Calculate volume
		if height[l] < height[r] {
			vol += max(0, mxhl-height[l])
			l++
			fmt.Println("volume:", vol, "mxhl:", mxhl, "l:", l)
		} else {
			vol += max(0, mxhr-height[r])
			r--
			fmt.Println("volume:", vol, "mxhr:", mxhr, "r:", r)
		}
	}
	return vol
}
	
