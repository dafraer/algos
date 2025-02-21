type stack struct {
	val  int
	ind  int
	next *stack
}

type Stack struct {
	s *stack
}

func (s *Stack) push(v, i int) {
	s.s = &stack{v, i, s.s}
}

func (s *Stack) pop() {
	s.s = s.s.next
}

func (s *Stack) top() (v, i int) {
	v = s.s.val
	i = s.s.ind
	return
}

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	s := &Stack{}
	for i, v := range temperatures {
		for s.s != nil {
			val, ind := s.top()
			if val < v {
				s.pop()
				ans[ind] = i - ind
			} else {
				s.push(v, i)
				break
			}
		}
		if s.s == nil {
			s.push(v, i)
		}
	}
	return ans
}
