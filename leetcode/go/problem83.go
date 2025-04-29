/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	m := make(map[int]*ListNode)
	cur := head
	var prev *ListNode
	for cur != nil {
		if _, ok := m[cur.Val]; ok {
			prev.Next = cur.Next
		} else {
			m[cur.Val] = cur
			prev = cur
		}
		cur = cur.Next
	}
	return head
}
