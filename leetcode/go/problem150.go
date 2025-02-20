type stack struct {
	val  int
	next *stack
}

func (s *stack) push(x int) {
	*s = stack{x, &stack{s.val, s.next}}
}

func (s *stack) pop() int {
	v := s.val
	*s = *s.next
	return v
}

func evalRPN(tokens []string) int {
	s := &stack{}
	for i := 0; i < len(tokens); i++ {
		if x, err := strconv.Atoi(tokens[i]); err == nil {
			s.push(x)
		} else {
			y := s.pop()
			x := s.pop()
			switch tokens[i] {
			case "+":
				s.push(x + y)
			case "-":
				s.push(x - y)
			case "*":
				s.push(x * y)
			case "/":
				s.push(x / y)
			}
		}
	}
	return s.pop()
}
