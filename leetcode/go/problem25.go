/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func pushFront(n *ListNode, val int) *ListNode {
	return &ListNode{Val: val, Next: n}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := head
	var ans *ListNode
	first := true
	var ad *ListNode
	for cur != nil {
		var rev *ListNode
		save := cur
		cnt := k
		var saveRev *ListNode
		for cnt > 0 {
			if cur == nil {
				ad.Next = save
				return ans
			}
			rev = pushFront(rev, cur.Val)
			if cnt == k && first {
				ad = rev
			}
			if cnt == k {
				saveRev = rev
			}
			cnt--
			cur = cur.Next
		}
		if first {
			ans = rev
			fmt.Println(ad.Val)
		} else {
			ad.Next = rev
			ad = saveRev
		}
		first = false
	}
	return ans
}
