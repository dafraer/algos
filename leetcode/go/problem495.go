func findPoisonedDuration(timeSeries []int, duration int) int {
	ans := duration
	for i := 0; i < len(timeSeries)-1; i++ {
		ans += min(duration, timeSeries[i+1]-timeSeries[i])
	}
	return ans
} 