/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := &ListNode{}
	cur := ans
	move := 0
	sum := 0
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			sum = l1.Val + l2.Val + move
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil {
			sum = l1.Val + move
			l1 = l1.Next
		} else {
			sum = l2.Val + move
			l2 = l2.Next
		}
		move = (sum - (sum % 10)) / 10
		cur.Val = sum % 10
		if l1 != nil || l2 != nil {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}
	if move != 0 {
		cur.Next = &ListNode{move, nil}
	}
	return ans
}
