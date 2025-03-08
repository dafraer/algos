/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	h := head.Next
	rev := head
	ln := 0
	for h != nil {
		rev = &ListNode{h.Val, &ListNode{rev.Val, rev.Next}}
		h = h.Next
		ln++
	}
	ans := head
	cnt := 1
	h = head.Next
	for ln > 0 {
		if cnt%2 != 0 {
			ans.Next = &ListNode{rev.Val, nil}
			rev = rev.Next
		} else {
			ans.Next = &ListNode{h.Val, nil}
			h = h.Next
		}
		ans = ans.Next
		cnt++
		ln--
	}
	ans = nil
}

