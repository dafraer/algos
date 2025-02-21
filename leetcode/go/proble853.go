type car struct {
	pos int
	v   int
}

type stack struct {
	c    car
	next *stack
}

type Stack struct {
	s *stack
}

func (s *Stack) push(c car) {
	s.s = &stack{c, s.s}
}

func (s *Stack) pop() {
	s.s = s.s.next
}

func (s *Stack) top() car {
	return s.s.c
}

func carFleet(target int, position []int, speed []int) int {
	a := make([]car, len(position))
	for i := 0; i < len(position); i++ {
		a[i].pos = position[i]
		a[i].v = speed[i]
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].pos < a[j].pos
	})
	cnt := 0
	s := &Stack{}
	for i := len(a) - 1; i >= 0; i-- {
		if i < len(a)-1 {
			topCar := s.top()
			if !doColide(topCar, a[i], target) {
				s.push(a[i])
				cnt++
			}
		} else {
			s.push(a[i])
			cnt++
		}
	}
	return cnt
}

func doColide(car1, car2 car, target int) bool {
	return float64((target-car2.pos))/float64(car2.v) <= float64((target-car1.pos))/float64(car1.v)
}
