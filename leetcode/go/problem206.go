/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var ans *ListNode
	for head != nil {
		ans = &ListNode{head.Val, ans}
		head = head.Next
	}
	return ans
}
