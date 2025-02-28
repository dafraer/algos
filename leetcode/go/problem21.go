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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	ans := &ListNode{}
	for list1 != nil || list2 != nil {
		if list1 == nil || (list2 != nil && list2.Val <= list1.Val) {
			ans = &ListNode{list2.Val, &ListNode{ans.Val, ans.Next}}
			list2 = list2.Next
		} else if list2 == nil || list1.Val <= list2.Val {
			ans = &ListNode{list1.Val, &ListNode{ans.Val, ans.Next}}
			list1 = list1.Next
		}
	}
	answer := reverseList(ans)
	return answer.Next
}
