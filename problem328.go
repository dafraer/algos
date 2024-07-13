func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := 0
	newHead := head
	for newHead != nil {
		n++
		newHead = newHead.Next
	}
	i := 0
    odd :=  &ListNode{}
	oddret := odd
	even := &ListNode{}
	evenret := even
	for head != nil {
		i++
		if i%2 == 0 {
			temp := &ListNode{}
			even.Val = head.Val
            if i == n || i == n-1 {
                even.Next = nil
            } else {
                even.Next = temp
            }
            if i == 2 {
                evenret = even
            }     
			even = even.Next
		} else {
			temp := &ListNode{}
            odd.Val = head.Val
            odd.Next = temp
			odd = &ListNode{head.Val, temp}
            if i == 1 {
                oddret = odd
            }
			odd = odd.Next
		}
		head = head.Next
	}
	odd.Val = evenret.Val
    odd.Next = evenret.Next
	return oddret
}
