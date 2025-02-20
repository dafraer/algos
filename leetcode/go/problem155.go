type MinStack struct {
	value int
	mn    int
	next  *MinStack
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	localMin := min(val, this.mn)
	if this.next == nil {
		localMin = val
	}
	*this = MinStack{val, localMin, &MinStack{this.value, this.mn, this.next}}
}

func (this *MinStack) Pop() {
	*this = *this.next
}

func (this *MinStack) Top() int {
	return this.value
}

func (this *MinStack) GetMin() int {
	return this.mn
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
