type deque struct {
	d []pair
}

type pair struct {
	val int
	ind int
}

func newDeque(ln int) *deque {
	a := make([]pair, 0, ln)
	return &deque{a}
}

func (d *deque) pushBack(p pair) {
	d.d = append(d.d, p)
}

func (d *deque) popBack() {
	d.d = d.d[:len(d.d)-1]
}

func (d *deque) topBack() pair {
	if len(d.d) == 0 {
		return pair{ind: -1}
	}
	return d.d[len(d.d)-1]
}

func (d *deque) topFront() pair {
	if len(d.d) == 0 {
		return pair{ind: -1}
	}
	return d.d[0]
}

func (d *deque) popFront() {
	d.d = d.d[1:]
}

func maxSlidingWindow(nums []int, k int) []int {
	d := newDeque(len(nums))
	ans := make([]int, 0, len(nums)-k)
	l := 0
	for i := 0; i < len(nums); i++ {
		for d.topBack().val < nums[i] && d.topBack().ind != -1 {
			d.popBack()
		}
		d.pushBack(pair{nums[i], i})
		if i > k-2 {
			if d.topFront().ind < l && d.topFront().ind != -1 {
				d.popFront()
			}
			ans = append(ans, d.topFront().val)
			l++
		}
	}
	return ans
}
