/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	for {
		if head == nil {
			break
		}
		if head.Val == math.MaxInt {
			return true
		}
		head.Val = math.MaxInt
		head = head.Next
	}
	return false
}
