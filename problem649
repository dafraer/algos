func predictPartyVictory(senate string) string {
	var qr, qd queue
	for i := 0; i < len(senate); i++ {
		if senate[i] == 'R' {
			qr.Push(i)
		} else {
			qd.Push(i)
		}
	}
	for len(qd) > 0 && len(qr) > 0 {
		if qr[0] > qd[0] {
			qr.Pop()
			qd.Push(qd[0] + len(senate))
			qd.Pop()
		} else {
			qd.Pop()
			qr.Push(qr[0] + len(senate))
			qr.Pop()
		}
	}
	if len(qd) > len(qr) {
		return "Dire"
	} else {
		return "Radiant"
	}
}

type queue []int

func (q *queue) Push(n int) {
	*q = append(*q, n)
}

func (q *queue) Pop() int {
	s := *q
	n := s[0]
	*q = s[1:]
	return n
}
