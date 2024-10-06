func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		var ans *ListNode
		return ans
	}
	n := 0
	newHead := head
	for newHead != nil {
		n++
		newHead = newHead.Next
	}
	i := 0
	ans := head
	for head != nil {
		i++
		if i == n/2 {
			head.Next = head.Next.Next
		}
		head = head.Next
	}
	return ans
}
