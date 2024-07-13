type RecentCounter struct {
	Requests []string
	Time     []int
}

func Constructor() RecentCounter {
	requests := make([]string, 0)
	requests = append(requests, "RecentCounter")
	time := make([]int, 1)
	return RecentCounter{
		requests,
		time,
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.Requests = append(this.Requests, "ping")
	this.Time = append(this.Time, t)
	cnt := 0
	for i := len(this.Time) - 1; i >= 1; i-- {
		if this.Time[len(this.Time)-1]-this.Time[i] > 3000 {
			break
		}
		cnt++
	}
	return cnt
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
