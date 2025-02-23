type pair struct {
	ind int
	val int
}

type stack struct {
	s []pair
}

func (s *stack) push(i, v int) {
	s.s = append(s.s, pair{i, v})
}
func (s *stack) pop() {
	if len(s.s) == 0 {
		panic("pop() called on an empty stack")
	}
	s.s = s.s[:len(s.s)-1]
}

func (s *stack) top() (int, int) {
	if len(s.s) == 0 {
		panic("top() called on an empty stack")
	}
	p := s.s[len(s.s)-1]
	return p.ind, p.val
}

func largestRectangleArea(heights []int) int {
	mx := 0
	s := &stack{}
	ind, val := 0, math.MaxInt
	for i, v := range heights {
		start := i
		if len(s.s) > 0 {
			ind, val = s.top()
		}
		for len(s.s) > 0 && val > v {
			mx = max(mx, val*(i-ind))
			start = ind
			s.pop()
			if len(s.s) <= 0 {
				break
			}
			ind, val = s.top()
		}
		s.push(start, v)
	}
	for len(s.s) > 0 {
		i, v := s.top()
		s.pop()
		mx = max(mx, v*(len(heights)-i))
	}
	return mx
}