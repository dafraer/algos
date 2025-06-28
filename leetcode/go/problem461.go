	func hammingDistance(x int, y int) int {
		var ans float64
		for x+y > 0 {
			ans += math.Abs(float64(x%2 - y%2))
			x /= 2
			y /= 2
		}
		return int(ans)
	}
