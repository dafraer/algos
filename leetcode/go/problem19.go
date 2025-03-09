/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	h := head
	length := 0
	for h != nil {
		length++
		h = h.Next
	}
	if length == n {
		return head.Next
	}
	cnt := 1
	h = head
	for h != nil {
		å
		if length-n == cnt {
			temp := h
			temp = temp.Next
			h.Next = temp.Next
			break
		}
		cnt++
		h = h.Next
	}
	return head
}
